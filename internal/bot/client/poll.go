package client

import (
	"context"

	"github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (i *ImapConcentratorClientImpl) PollV1(ctx context.Context, chatID int64) error {
	_, err := i.client.PollV1(ctx, &imap_concentrator.PollV1Request{ChatID: chatID})
	if err != nil {
		return err
	}

	return nil
}
