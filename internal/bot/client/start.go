package client

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (i *ImapConcentratorClientImpl) BeginConversation(ctx context.Context, chatID int64) error {
	_, err := i.client.InsertChatV1(ctx, &pb.InsertChatV1Request{
		ChatID: chatID,
	})

	return err
}
