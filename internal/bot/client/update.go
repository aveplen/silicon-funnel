package client

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (i *ImapConcentratorClientImpl) UpdateMailbox(ctx context.Context, chatID int64, mailbox *pb.MailboxV1) error {

	_, err := i.client.UpdateMailboxV1(ctx, &pb.UpdateMailboxV1Request{
		ChatID:  chatID,
		Mailbox: mailbox,
	})

	return err
}
