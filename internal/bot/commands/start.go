package commands

import (
	"github.com/aveplen/silicon-funnel/internal/bot/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _ ServiceStartCapabilities = (*service.Service)(nil)

type ServiceStartCapabilities interface {
	Start(chatID int64) error
}

func (c *CommandHandler) Start(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	if err := c.service.Start(chatID); err != nil {
		panic(err)
	}
}
