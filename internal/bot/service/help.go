package service

import "github.com/aveplen/silicon-funnel/internal/bot/telegram"

var _ TelegramHelpCapabilities = (*telegram.TelegramService)(nil)

type TelegramHelpCapabilities interface {
	SendHelpText(chatID int64) error
}

func (s *Service) Help(chatID int64) error {
	return s.telegramService.SendHelpText(chatID)
}
