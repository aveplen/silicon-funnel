package commands

import (
	"github.com/aveplen/silicon-funnel/internal/bot/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _ ServiceAcquireCapabilities = (*service.Service)(nil)

type ServiceAcquireCapabilities interface {
	Acquire(chatID int64) error
}

func (c *CommandHandler) Acquire(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	if err := c.service.Acquire(chatID); err != nil {
		panic(err)
	}
}
