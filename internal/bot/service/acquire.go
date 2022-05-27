package service

import (
	"github.com/aveplen/silicon-funnel/internal/bot/telegram"
)

var _ TelegramAcquireCapabilities = (*telegram.TelegramService)(nil)

type TelegramAcquireCapabilities interface {
	ResponseToAcquire(chatID int64) error
}

func (s *Service) Acquire(chatID int64) error {
	return s.telegramService.ResponseToAcquire(chatID)
}
