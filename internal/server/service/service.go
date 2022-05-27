package service

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/aveplen/silicon-funnel/internal/server/client"
	"github.com/aveplen/silicon-funnel/internal/server/repo"
	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/sirupsen/logrus"
)

type subscription struct {
	clientID int64
	stream   pb.ImapConcentrator_SubscribeToAllNotificationsV1Server
	finished chan<- bool
}

type allSubscribers struct {
	allSubscribers []*subscription
	mu             *sync.Mutex
}

type chatSubscribers struct {
	chatSubscribers map[int64][]*subscription
	mu              *sync.Mutex
}

type ImapConcentratorServer struct {
	pb.UnimplementedImapConcentratorServer
	whitelist      []string
	notificationch chan *pb.NotificationV1
	clientIDSec    *atomic.Value
	repo           *repo.Repository
	client         *client.ImapClient
	allSubs        *allSubscribers
	chatSubs       *chatSubscribers
}

func NewImapConcentratorServer(
	whitelist []string,
	repo *repo.Repository,
	client *client.ImapClient,
	notificationch chan *pb.NotificationV1,
) (*ImapConcentratorServer, error) {

	clientIDSeq := &atomic.Value{}
	clientIDSeq.Store(int64(1))

	return &ImapConcentratorServer{
		whitelist:      whitelist,
		notificationch: notificationch,
		clientIDSec:    clientIDSeq,
		repo:           repo,
		client:         client,
		allSubs: &allSubscribers{
			allSubscribers: []*subscription{},
			mu:             &sync.Mutex{},
		},
		chatSubs: &chatSubscribers{
			chatSubscribers: map[int64][]*subscription{},
			mu:              &sync.Mutex{},
		},
	}, nil
}

func (i *ImapConcentratorServer) ListenForSubscribers(ctx context.Context) error {
	logrus.Infof("listening for subscribers")

	for {
		var disconnected []*subscription

		select {
		case notification := <-i.notificationch:
			subs := i.chatSubs.chatSubscribers[notification.ChatID]
			subs = append(subs, i.allSubs.allSubscribers...)

			for _, sub := range subs {
				if err := sub.stream.Send(notification); err != nil {
					disconnected = append(disconnected, sub)
				}
			}

			for _, dis := range disconnected {
				dis.finished <- true
			}

		case <-ctx.Done():
			return fmt.Errorf("context done")
		}
	}
}
