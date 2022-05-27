package service

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ImapConcentratorServer) InsertChatV1(ctx context.Context, req *pb.InsertChatV1Request) (*pb.Ack, error) {

	logrus.Infof("InsertChatV1 called")

	if err := i.repo.InsertChat(ctx, req.ChatID); err != nil {
		return nil, status.Errorf(codes.Unavailable, "%w", err)
	}

	return &pb.Ack{}, nil
}
