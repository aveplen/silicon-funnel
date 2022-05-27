package service

import (
	"context"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/sirupsen/logrus"
)

func (i *ImapConcentratorServer) UnsubscribeFromNotificationV1(
	ctx context.Context,
	req *pb.UnsubscribeFromNotificationsV1Request,
) (*pb.Ack, error) {

	logrus.Infof("UnsubscribeFromNotificationV1 called")

	func() {
		i.allSubs.mu.Lock()
		defer i.allSubs.mu.Unlock()

		allSubs := i.allSubs.allSubscribers
		for i, sub := range allSubs {
			if sub.clientID == req.ClientID {
				allSubs[0], allSubs[i] = allSubs[i], allSubs[0]
				allSubs = allSubs[1:]
			}
		}
	}()

	func() {
		i.chatSubs.mu.Lock()
		defer i.chatSubs.mu.Unlock()

		chatSubs, ok := i.chatSubs.chatSubscribers[req.ClientID]
		if !ok {
			return
		}
		for i, sub := range chatSubs {
			if sub.clientID == req.ClientID {
				chatSubs[0], chatSubs[i] = chatSubs[i], chatSubs[0]
				chatSubs = chatSubs[1:]
			}
		}
	}()

	return &pb.Ack{}, nil
}
