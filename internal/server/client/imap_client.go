package client

import (
	"context"
	"fmt"

	"github.com/aveplen/silicon-funnel/internal/server/repo"
	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ImapClient struct {
	db            *repo.Repository
	notifications chan *pb.NotificationV1
	errs          chan error
}

func NewImapClient(
	db *repo.Repository,
	notifications chan *pb.NotificationV1,
	errs chan error,
) (*ImapClient, error) {

	return &ImapClient{
		db:            db,
		notifications: notifications,
		errs:          errs,
	}, nil
}

func (i *ImapClient) PollFor(ctx context.Context, tgChatID int64) error {
	logrus.Infof("polling mailboxes for %v", tgChatID)

	mailboxes, metas, err := i.db.GetMailboxesByTgChatID(ctx, tgChatID)
	if err != nil {
		return err
	}

	return i.pollMailboxes(ctx, mailboxes, metas)
}

func (i *ImapClient) Poll(ctx context.Context) error {
	logrus.Infof("polling all mailboxes")

	mailboxes, metas, err := i.db.GetAllMailboxes(ctx)
	if err != nil {
		return err
	}

	return i.pollMailboxes(ctx, mailboxes, metas)
}

func (i *ImapClient) pollMailboxes(ctx context.Context, mailboxes []*pb.MailboxV1, metas []repo.MailboxMeta) error {
	for j, mailbox := range mailboxes {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context done")
		default:
		}

		c, err := client.DialTLS(fmt.Sprintf("%s:%d", mailbox.Host, mailbox.Port), nil)
		if err != nil {
			return err
		}

		if err := c.Login(mailbox.Username, mailbox.Password); err != nil {
			return err
		}

		mbox, err := c.Select(mailbox.Mailbox, true)
		if err != nil {
			return err
		}

		seqset := new(imap.SeqSet)
		seqset.AddRange(metas[j].Offset+1, mbox.Messages+1)

		i.db.PostBump(&repo.OffsetBump{
			MailboxID: mailbox.MailboxID,
			Bump:      int(mbox.Messages - metas[j].Offset),
		})

		messages := make(chan *imap.Message, 100)
		go func() {
			if err := c.Fetch(seqset, []imap.FetchItem{imap.FetchFull}, messages); err != nil {
				i.errs <- err
			}
		}()

		for msg := range messages {
			i.notifications <- &pb.NotificationV1{
				ChatID:  metas[j].TgChatID,
				Sender:  msg.Envelope.Sender[0].MailboxName,
				Email:   msg.Envelope.To[0].MailboxName,
				Host:    mailbox.Host,
				Port:    mailbox.Port,
				Mailbox: mailbox.Mailbox,
				Title:   msg.Envelope.Subject,
				Timestamp: &timestamppb.Timestamp{
					Seconds: msg.Envelope.Date.Unix(),
				},
			}
		}
	}
	return nil
}
