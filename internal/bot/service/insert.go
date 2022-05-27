package service

import (
	"fmt"

	"github.com/aveplen/silicon-funnel/internal/bot/builder"
	"github.com/aveplen/silicon-funnel/internal/bot/telegram"
	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

var _ MailboxBuilderServiceInsertCapabilities = (*builder.MailboxBuilderService)(nil)

type MailboxBuilderServiceInsertCapabilities interface {
	Mount(chatID int64, builder *builder.MailboxBuilder) error
	Get(chatID int64) (*builder.MailboxBuilder, bool)
}

var _ TelegramInsertCapabilities = (*telegram.TelegramService)(nil)

type TelegramInsertCapabilities interface {
	AskForHost(chatID int64) error
	AskForPort(chatID int64) error
	AskForMailbox(chatID int64) error
	AskForUsername(chatID int64) error
	AskForPassword(chatID int64) error
	AskForSubmit(chatID int64, request *pb.MailboxV1) error
}

func (s *Service) Insert(chatID int64) error {
	_, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		mb := builder.NewMailboxBuilder()
		s.mailboxBuilderService.Mount(chatID, mb)
	}
	return s.telegramService.AskForHost(chatID)
}

func (s *Service) BuilderMounted(chatID int64) bool {
	_, ok := s.mailboxBuilderService.Get(chatID)
	return ok
}

func (s *Service) GetBuilderStage(chatID int64) (int, error) {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return 0, fmt.Errorf("builder not mounted on %d", chatID)
	}

	return b.Stage, nil
}

func (s *Service) Submit(chatID int64) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	return s.telegramService.AskForSubmit(chatID, b.Mailbox)
}

func (s *Service) Initial(chatID int64) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Stage = builder.Host
	return s.telegramService.AskForHost(chatID)
}

func (s *Service) Host(chatID int64, host string) error {
	fmt.Println("host")

	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Mailbox.Host = host

	b.Stage = builder.Port
	return s.telegramService.AskForPort(chatID)
}

func (s *Service) Port(chatID int64, port int32) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Mailbox.Port = port

	b.Stage = builder.Mailbox
	return s.telegramService.AskForMailbox(chatID)
}

func (s *Service) Mailbox(chatID int64, mailbox string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Mailbox.Mailbox = mailbox

	b.Stage = builder.Username
	return s.telegramService.AskForUsername(chatID)
}

func (s *Service) Username(chatID int64, username string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Mailbox.Username = username

	b.Stage = builder.Password
	return s.telegramService.AskForPassword(chatID)
}

func (s *Service) Password(chatID int64, password string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Mailbox.Password = password

	b.Stage = builder.Submit
	return s.telegramService.AskForSubmit(chatID, b.Mailbox)
}

func (s *Service) CorrectHost(chatID int64, host string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Mailbox.Host = host

	b.Stage = builder.Submit
	return s.telegramService.AskForSubmit(chatID, b.Mailbox)
}

func (s *Service) CorrectPort(chatID int64, port int32) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Mailbox.Port = port

	b.Stage = builder.Submit
	return s.telegramService.AskForSubmit(chatID, b.Mailbox)
}

func (s *Service) CorrectMailbox(chatID int64, mailbox string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Mailbox.Mailbox = mailbox

	b.Stage = builder.Submit
	return s.telegramService.AskForSubmit(chatID, b.Mailbox)
}

func (s *Service) CorrectUsername(chatID int64, username string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Mailbox.Username = username

	b.Stage = builder.Submit
	return s.telegramService.AskForSubmit(chatID, b.Mailbox)
}

func (s *Service) CorrectPassword(chatID int64, password string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Mailbox.Password = password

	b.Stage = builder.Submit
	return s.telegramService.AskForSubmit(chatID, b.Mailbox)
}
