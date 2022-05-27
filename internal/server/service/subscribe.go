package service

import (
	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (i *ImapConcentratorServer) SubscribeToAllNotificationsV1(
	req *pb.SubscribeToAllNotificationsV1Request,
	stream pb.ImapConcentrator_SubscribeToAllNotificationsV1Server,
) error {

	logrus.Infof("SubscribeToAllNotificationsV1 called")

	if !i.whitelisted(req.Key) {
		return status.Errorf(codes.Unauthenticated, "key is not whitelisted: <%v>", req.Key)
	}

	finished := make(chan bool)

	if err := func() error {
		i.allSubs.mu.Lock()
		defer i.allSubs.mu.Unlock()

		for _, c := range i.allSubs.allSubscribers {
			if req.ClientID == c.clientID {
				return status.Errorf(codes.Unavailable, "client id already taken: <%d>", req.ClientID)
			}
		}

		i.allSubs.allSubscribers = append(i.allSubs.allSubscribers, &subscription{
			clientID: req.ClientID,
			stream:   stream,
			finished: finished,
		})

		return nil

	}(); err != nil {
		return err
	}

	select {
	case <-finished:
		return status.Errorf(codes.Unavailable, "unexpected client otval: %d", req.ClientID)

	case <-stream.Context().Done():
		return status.Errorf(codes.Aborted, "context done: %d", req.ClientID)
	}
}

func (i *ImapConcentratorServer) SubscribeToChatNotificationsV1(
	req *pb.SubscribeToChatNotificationsV1Request,
	stream pb.ImapConcentrator_SubscribeToChatNotificationsV1Server,
) error {

	finished := make(chan bool)

	if err := func() error {
		i.chatSubs.mu.Lock()
		defer i.chatSubs.mu.Unlock()

		chatSubs, ok := i.chatSubs.chatSubscribers[req.ChatID]
		if !ok {
			i.chatSubs.chatSubscribers[req.ChatID] = []*subscription{}
			chatSubs = i.chatSubs.chatSubscribers[req.ChatID]
		}

		for _, c := range chatSubs {
			if req.ClientID == c.clientID {
				return status.Errorf(codes.Unavailable, "client id already taken: %d", req.ClientID)
			}
		}

		i.chatSubs.chatSubscribers[req.ChatID] = append(i.chatSubs.chatSubscribers[req.ChatID], &subscription{
			clientID: req.ClientID,
			stream:   stream,
			finished: finished,
		})

		return nil

	}(); err != nil {
		return err
	}

	select {
	case <-finished:
		return status.Errorf(codes.Unavailable, "unexpected client otval: %d", req.ClientID)

	case <-stream.Context().Done():
		return status.Errorf(codes.Aborted, "context done: %d", req.ClientID)
	}
}

func (i *ImapConcentratorServer) whitelisted(key string) bool {
	for _, whkey := range i.whitelist {
		if whkey == key {
			return true
		}
	}
	return false
}
