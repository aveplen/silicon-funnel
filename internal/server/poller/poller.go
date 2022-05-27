package poller

import (
	"context"

	"github.com/aveplen/silicon-funnel/internal/server/client"
	"github.com/aveplen/silicon-funnel/internal/server/config"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type Poller struct {
	cronInst *cron.Cron
	cronErrs chan error
	client   *client.ImapClient
}

func NewPoller(cfg *config.Config, client *client.ImapClient) (*Poller, error) {
	poller := &Poller{
		cronInst: cron.New(),
		cronErrs: make(chan error, 100),
		client:   client,
	}

	if _, err := poller.cronInst.AddFunc(cfg.Poller.Cron, func() {
		if err := poller.client.Poll(context.Background()); err != nil {
			poller.cronErrs <- err
		}
	}); err != nil {
		return nil, err
	}

	return poller, nil
}

func (p *Poller) Start(ctx context.Context) {
	logrus.Infof("starting poller")

	p.cronInst.Start()
	<-ctx.Done()

	jobs := p.cronInst.Stop()
	<-jobs.Done()
}
