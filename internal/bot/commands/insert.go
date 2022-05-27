package commands

import (
	"github.com/aveplen/silicon-funnel/internal/bot/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _ ServiceInsertCapabilities = (*service.Service)(nil)

type ServiceInsertCapabilities interface {
	Insert(chatID int64) error
}

func (c *CommandHandler) Insert(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	if err := c.service.Insert(chatID); err != nil {
		panic(err)
	}
}
