package builder

import (
	"fmt"
	"sync"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

// ================================== builder =================================

const (
	Initial = iota

	Host
	Port
	Mailbox
	Username
	Password

	CorrectHost
	CorrectPort
	CorrectMailbox
	CorrectUsername
	CorrectPassword

	Submit
)

type MailboxBuilder struct {
	Stage   int
	Mailbox *pb.MailboxV1
}

func NewMailboxBuilder() *MailboxBuilder {
	return &MailboxBuilder{
		Stage:   Host,
		Mailbox: &pb.MailboxV1{},
	}
}

// ================================== service =================================
type MailboxBuilderService struct {
	mounted map[int64]*MailboxBuilder
	mu      *sync.Mutex
}

func NewMailboxBuilderService() *MailboxBuilderService {
	return &MailboxBuilderService{
		mounted: make(map[int64]*MailboxBuilder),
		mu:      &sync.Mutex{},
	}
}

func (b *MailboxBuilderService) Mount(chatID int64, builder *MailboxBuilder) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.mounted[chatID]; ok {
		return fmt.Errorf("another builder is already mounted on %d", chatID)
	}

	b.mounted[chatID] = builder
	return nil
}

func (b *MailboxBuilderService) Unmount(chatID int64) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.mounted[chatID]; !ok {
		return fmt.Errorf("no builder is mounted on %d", chatID)
	}

	delete(b.mounted, chatID)
	return nil
}

func (b *MailboxBuilderService) Get(chatID int64) (*MailboxBuilder, bool) {
	b.mu.Lock()
	defer b.mu.Unlock()

	builder, ok := b.mounted[chatID]
	return builder, ok
}
