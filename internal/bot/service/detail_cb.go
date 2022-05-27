package service

import (
	"context"

	"github.com/aveplen/silicon-funnel/internal/bot/telegram"
	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/sirupsen/logrus"
)

var _ TelegramDetailCallbackCapabilities = (*telegram.TelegramService)(nil)

type TelegramDetailCallbackCapabilities interface {
	CloseKeyboard(chatID int64, messageID int, text string) error
	ReturnToMailboxes(chatID int64, messageID int, mailboxes []*pb.MailboxV1) error
}

func (s *Service) ReturnToList(chatID int64, mailboxID int, messageID int) error {
	logrus.Info("ReturnToList")

	mailboxes, err := s.client.ListMailboxesV1(context.Background(), chatID)
	if err != nil {
		return err
	}

	return s.telegramService.ReturnToMailboxes(chatID, messageID, mailboxes)
}

func (s *Service) DeleteMailbox(chatID int64, mailboxID int64, messageID int) error {
	logrus.Info("DeleteMailbox")
	// return s.client.DeleteMailbox(context.Background(), chatID, mailboxID)
	return s.telegramService.CloseKeyboard(chatID, messageID, "Удалено")
}
