package service

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ImapConcentratorServer) DeleteMailboxV1(ctx context.Context, req *pb.DeleteMailboxV1Request) (*pb.Ack, error) {

	logrus.Infof("DeleteMailboxV1 called")

	err := i.repo.DeleteMailbox(ctx, req.MailboxID)
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	return &pb.Ack{}, nil
}
