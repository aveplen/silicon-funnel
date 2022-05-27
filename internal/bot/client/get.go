package client

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (i *ImapConcentratorClientImpl) GetMailboxV1(
	ctx context.Context,
	chatID int64,
	mailboxID int64,
) (*pb.MailboxV1, error) {

	res, err := i.client.GetMailboxV1(ctx, &pb.GetMailboxV1Request{
		ChatID:    chatID,
		MailboxID: mailboxID,
	})

	if err != nil {
		return nil, err
	}

	return res.Mailbox, nil
}
