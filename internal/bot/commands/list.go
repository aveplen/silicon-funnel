package commands

import (
	"github.com/aveplen/silicon-funnel/internal/bot/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _ ServiceListCapabilities = (*service.Service)(nil)

type ServiceListCapabilities interface {
	ListMailboxes(chatID int64) error
}

func (c *CommandHandler) List(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID

	if err := c.service.ListMailboxes(chatID); err != nil {
		panic(err)
	}
}
