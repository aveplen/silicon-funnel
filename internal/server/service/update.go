package service

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ImapConcentratorServer) UpdateMailboxV1(ctx context.Context, req *pb.UpdateMailboxV1Request) (*pb.Ack, error) {

	logrus.Infof("UpdateMailboxV1 called")

	if err := i.repo.UpdateMailbox(ctx, req.Mailbox); err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	return &pb.Ack{}, nil
}
