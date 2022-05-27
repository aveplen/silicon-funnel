package service

import (
	"context"

	"github.com/aveplen/silicon-funnel/internal/server/client"
	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ImapConcentratorServer) InsertEmailV1(ctx context.Context, req *pb.InsertMailboxV1Request) (*pb.Ack, error) {

	logrus.Infof("InsertEmailV1 called")

	offset, err := client.GetOffset(ctx, req.Mailbox)
	if err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	if err := i.repo.InsertChat(ctx, req.ChatID); err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	if err := i.repo.InsertMailbox(ctx, req.ChatID, req.Mailbox, offset); err != nil {
		return nil, status.Error(codes.Unavailable, err.Error())
	}

	return &pb.Ack{}, nil
}
