package commands

import (
	"github.com/aveplen/silicon-funnel/internal/bot/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _ ServiceHelpCapabilities = (*service.Service)(nil)

type ServiceHelpCapabilities interface {
	Help(chatID int64) error
}

func (c *CommandHandler) Help(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	if err := c.service.Help(chatID); err != nil {
		panic(err)
	}
}
