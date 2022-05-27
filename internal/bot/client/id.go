package client

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

func (i *ImapConcentratorClientImpl) AssignUniqueClientID(ctx context.Context) (int64, error) {
	res, err := i.client.AssignUniqueClientID(ctx, &pb.AssignUniqueClientIDRequest{})

	if err != nil {
		return 0, err
	}

	return res.ClientID, nil
}
