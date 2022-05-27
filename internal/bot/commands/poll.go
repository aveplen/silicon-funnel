package commands

import (
	"github.com/aveplen/silicon-funnel/internal/bot/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _ ServicePollCapabilities = (*service.Service)(nil)

type ServicePollCapabilities interface {
	Poll(chatID int64) error
}

func (c *CommandHandler) Poll(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	if err := c.service.Poll(chatID); err != nil {
		panic(err)
	}
}
