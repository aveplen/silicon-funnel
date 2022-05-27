package service

import (
	"context"

	"github.com/aveplen/silicon-funnel/internal/bot/telegram"
	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
)

var _ TelegramListCapabilities = (*telegram.TelegramService)(nil)

type TelegramListCapabilities interface {
	SendMailboxesList(chatID int64, mailboxes []*pb.MailboxV1) error
}

func (s *Service) ListMailboxes(chatID int64) error {
	list, err := s.client.ListMailboxesV1(context.Background(), chatID)
	if err != nil {
		return err
	}

	return s.telegramService.SendMailboxesList(chatID, list)
}
