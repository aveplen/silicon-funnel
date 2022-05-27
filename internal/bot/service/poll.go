package service

import (
	"context"

	"github.com/aveplen/silicon-funnel/internal/bot/telegram"
)

var _ TelegramPollCapabilities = (*telegram.TelegramService)(nil)

type TelegramPollCapabilities interface {
}

func (s *Service) Poll(chatID int64) error {
	return s.client.PollV1(context.Background(), chatID)
}
