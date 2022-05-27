package service

import (
	"context"

	"github.com/aveplen/silicon-funnel/internal/bot/telegram"
)

var _ TelegramStartCapabilities = (*telegram.TelegramService)(nil)

type TelegramStartCapabilities interface {
	NotifyAboutStart(chatID int64) error
}

func (s *Service) Start(chatID int64) error {
	if err := s.client.BeginConversation(context.Background(), chatID); err != nil {
		return err
	}
	return s.telegramService.NotifyAboutStart(chatID)
}
