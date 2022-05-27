package service

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ImapConcentratorServer) ListMailboxesV1(
	ctx context.Context,
	req *pb.ListMailboxesV1Request,
) (*pb.ListMailboxesV1Response, error) {

	logrus.Infof("ListMailboxesV1 called")

	mailboxes, _, err := i.repo.GetMailboxesByTgChatID(ctx, req.ChatID)
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	return &pb.ListMailboxesV1Response{Mailboxes: mailboxes}, nil
}
