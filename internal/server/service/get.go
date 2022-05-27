package service

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ImapConcentratorServer) GetMailboxV1(
	ctx context.Context,
	req *pb.GetMailboxV1Request,
) (*pb.GetMailboxV1Response, error) {

	logrus.Infof("GetMailboxV1 called")

	mailbox, err := i.repo.GetMailboxByID(ctx, req.MailboxID)
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	return &pb.GetMailboxV1Response{
		ChatID:  req.ChatID,
		Mailbox: mailbox,
	}, nil
}
