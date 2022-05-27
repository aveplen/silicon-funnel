package service

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ImapConcentratorServer) PollV1(ctx context.Context, req *pb.PollV1Request) (*pb.Ack, error) {

	logrus.Infof("PollV1 called")

	if err := i.client.PollFor(ctx, req.ChatID); err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	return &pb.Ack{}, nil
}
