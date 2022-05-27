package client

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (i *ImapConcentratorClientImpl) UnsubscribeFromNotifications(ctx context.Context, clientID int64) error {
	_, err := i.client.UnsubscribeFromNotificationsV1(ctx, &pb.UnsubscribeFromNotificationsV1Request{
		ClientID: clientID,
	})

	return err
}

func (i *ImapConcentratorClientImpl) unsubscribe(ctx context.Context) error {
	return i.UnsubscribeFromNotifications(ctx, i.clientID)
}
