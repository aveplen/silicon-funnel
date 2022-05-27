package client

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (i *ImapConcentratorClientImpl) ListMailboxesV1(ctx context.Context, chatID int64) ([]*pb.MailboxV1, error) {
	res, err := i.client.ListMailboxesV1(ctx, &pb.ListMailboxesV1Request{
		ChatID: chatID,
	})

	if err != nil {
		return nil, err
	}

	return res.Mailboxes, nil
}
