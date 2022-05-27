package client

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (i *ImapConcentratorClientImpl) DeleteMailbox(ctx context.Context, chatID int64, mailboxID int64) error {

	_, err := i.client.DeleteMailboxV1(ctx, &pb.DeleteMailboxV1Request{
		ChatID:    chatID,
		MailboxID: mailboxID,
	})

	return err
}
