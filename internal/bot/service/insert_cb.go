package service

import (
	"context"
	"fmt"

	"github.com/aveplen/silicon-funnel/internal/bot/builder"
	"github.com/aveplen/silicon-funnel/internal/bot/telegram"
	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

var _ MailboxBuilderServiceInsertCallbackCapabilities = (*builder.MailboxBuilderService)(nil)

type MailboxBuilderServiceInsertCallbackCapabilities interface {
	Unmount(chatID int64) error
	Get(chatID int64) (*builder.MailboxBuilder, bool)
}

var _ TelegramInsertCallbackCapabilities = (*telegram.TelegramService)(nil)

type TelegramInsertCallbackCapabilities interface {
	InsertSubmitResponse(chatID int64, messageID int) error
	AskForHost(chatID int64) error
	AskForPort(chatID int64) error
	AskForMailbox(chatID int64) error
	AskForUsername(chatID int64) error
	AskForPassword(chatID int64) error
	AskForSubmit(chatID int64, request *pb.MailboxV1) error
}

func (s *Service) SubmitCallback(chatID int64, messageID int, text string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	if err := s.client.InsertMailbox(context.Background(), chatID, b.Mailbox); err != nil {
		return err
	}

	if err := s.mailboxBuilderService.Unmount(chatID); err != nil {
		return err
	}

	return s.telegramService.InsertSubmitResponse(chatID, messageID)
}

func (s *Service) CorrectHostCallback(chatID int64, messageID int, text string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Stage = builder.CorrectHost
	if err := s.telegramService.AskForHost(chatID); err != nil {
		return err
	}

	return s.telegramService.CloseKeyboard(chatID, messageID, text)
}

func (s *Service) CorrectPortCallback(chatID int64, messageID int, text string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Stage = builder.CorrectPort
	if err := s.telegramService.AskForPort(chatID); err != nil {
		return err
	}

	return s.telegramService.CloseKeyboard(chatID, messageID, text)
}

func (s *Service) CorrectMailboxCallback(chatID int64, messageID int, text string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Stage = builder.CorrectMailbox
	if err := s.telegramService.AskForMailbox(chatID); err != nil {
		return err
	}

	return s.telegramService.CloseKeyboard(chatID, messageID, text)
}

func (s *Service) CorrectUsernameCallback(chatID int64, messageID int, text string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Stage = builder.CorrectUsername
	if err := s.telegramService.AskForUsername(chatID); err != nil {
		return err
	}

	return s.telegramService.CloseKeyboard(chatID, messageID, text)
}

func (s *Service) CorrectPasswordCallback(chatID int64, messageID int, text string) error {
	b, ok := s.mailboxBuilderService.Get(chatID)
	if !ok {
		return fmt.Errorf("builder not mounted on %d", chatID)
	}

	b.Stage = builder.CorrectPassword
	if err := s.telegramService.AskForPassword(chatID); err != nil {
		return err
	}

	return s.telegramService.CloseKeyboard(chatID, messageID, text)
}
