package client

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (i *ImapConcentratorClientImpl) InsertMailbox(ctx context.Context, chatID int64, mailbox *pb.MailboxV1) error {

	_, err := i.client.InsertMailboxV1(ctx, &pb.InsertMailboxV1Request{
		ChatID:  chatID,
		Mailbox: mailbox,
	})

	return err
}
