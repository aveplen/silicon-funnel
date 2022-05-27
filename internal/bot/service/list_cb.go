package service

import (
	"context"

	"github.com/aveplen/silicon-funnel/internal/bot/telegram"
	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

var _ TelegramListCallbackCapabilities = (*telegram.TelegramService)(nil)

type TelegramListCallbackCapabilities interface {
	ChangeKeyboardToDetail(chatID int64, messageID int, mailbox *pb.MailboxV1) error
}

func (s *Service) MailboxDetail(chatID int64, mailboxID int64, messageID int) error {
	mailbox, err := s.client.GetMailboxV1(context.Background(), chatID, mailboxID)
	if err != nil {
		return err
	}

	return s.telegramService.ChangeKeyboardToDetail(chatID, messageID, mailbox)
}
