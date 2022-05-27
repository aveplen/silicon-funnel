//go:generate mkdir -p mock
//go:generate minimock -o ./mock/ -s .go -g

package client

import (
	"context"
	"fmt"
	"log"

	"github.com/aveplen/silicon-funnel/internal/bot/config"
	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var _ ImapConcentratorClient = (*ImapConcentratorClientImpl)(nil)

type ImapConcentratorClient interface {
	DeleteMailbox(ctx context.Context, chatID int64, mailboxID int64) error
	AssignUniqueClientID(ctx context.Context) (int64, error)
	InsertMailbox(ctx context.Context, chatID int64, mailbox *pb.MailboxV1) error
	ListMailboxesV1(ctx context.Context, chatID int64) ([]*pb.MailboxV1, error)
	GetMailboxV1(ctx context.Context, chatID int64, mailboxID int64) (*pb.MailboxV1, error)
	PollV1(ctx context.Context, chatID int64) error
	BeginConversation(ctx context.Context, chatID int64) error
	SubscribeToAllNotifications(ctx context.Context, clientID int64) (
		pb.ImapConcentrator_SubscribeToAllNotificationsV1Client,
		error,
	)
	UnsubscribeFromNotifications(ctx context.Context, clientID int64) error
	UpdateMailbox(ctx context.Context, chatID int64, mailbox *pb.MailboxV1) error
}

type ImapConcentratorClientImpl struct {
	key           string
	client        pb.ImapConcentratorClient
	conn          *grpc.ClientConn
	notifications chan *pb.NotificationV1
	errors        chan error
	clientID      int64
}

func NewImapConcentratorClient(
	cfg *config.Config,
	notifications chan *pb.NotificationV1,
	errors chan error,
) (*ImapConcentratorClientImpl, error) {

	conn, err := grpc.Dial(cfg.Imapc.Addr, []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithReturnConnectionError()}...,
	)
	if err != nil {
		log.Printf("could not dial")
		return nil, err
	}

	return &ImapConcentratorClientImpl{
		key:           cfg.Imapc.Key,
		client:        pb.NewImapConcentratorClient(conn),
		conn:          conn,
		notifications: notifications,
		errors:        errors,
	}, nil
}

func (i *ImapConcentratorClientImpl) Start(ctx context.Context) error {
	stream, err := i.subscribe(ctx)
	if err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context done")
		default:
		}

		response, err := stream.Recv()
		if err != nil {
			i.errors <- err
			continue
		}
		i.notifications <- response
	}
}
