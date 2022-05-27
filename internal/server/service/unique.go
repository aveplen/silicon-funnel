package service

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ImapConcentratorServer) AssignUniqueClientID(
	context.Context,
	*pb.AssignUniqueClientIDRequest,
) (*pb.AssignUniqueClientIDResponse, error) {

	logrus.Infof("AssignUniqueClientID called")

	atom := i.clientIDSec.Load()
	clientID, ok := atom.(int64)
	if !ok {
		return nil, status.Error(codes.Internal, "type assertion error")
	}
	i.clientIDSec.Store(clientID + 1)

	return &pb.AssignUniqueClientIDResponse{ClientID: clientID}, nil
}
