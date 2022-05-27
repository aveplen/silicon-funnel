package client

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (i *ImapConcentratorClientImpl) SubscribeToAllNotifications(
	ctx context.Context,
	clientID int64,
) (pb.ImapConcentrator_SubscribeToAllNotificationsV1Client, error) {

	return i.client.SubscribeToAllNotificationsV1(ctx, &pb.SubscribeToAllNotificationsV1Request{
		Key:      i.key,
		ClientID: clientID,
	})
}

func (i *ImapConcentratorClientImpl) subscribe(ctx context.Context) (
	pb.ImapConcentrator_SubscribeToAllNotificationsV1Client,
	error,
) {

	clientID, err := i.AssignUniqueClientID(ctx)
	if err != nil {
		return nil, err
	}

	i.clientID = clientID

	return i.SubscribeToAllNotifications(ctx, i.clientID)
}
