package callback

import (
	"strings"

	"github.com/aveplen/silicon-funnel/internal/bot/service"
	"github.com/aveplen/silicon-funnel/internal/bot/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _ ServiceInsertCallbackCapabilities = (*service.Service)(nil)

type ServiceInsertCallbackCapabilities interface {
	SubmitCallback(chatID int64, messageID int, text string) error
	CorrectHostCallback(chatID int64, messageID int, text string) error
	CorrectPortCallback(chatID int64, messageID int, text string) error
	CorrectMailboxCallback(chatID int64, messageID int, text string) error
	CorrectUsernameCallback(chatID int64, messageID int, text string) error
	CorrectPasswordCallback(chatID int64, messageID int, text string) error
}

func (c *CallbackQueryHandler) Insert(update tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	action := strings.Split(update.CallbackData(), "/")[2]
	messageID := update.CallbackQuery.Message.MessageID
	text := update.CallbackQuery.Message.Text

	switch action {
	case telegram.InsertCorrectHostButtonData:
		if err := c.service.CorrectHostCallback(chatID, messageID, text); err != nil {
			panic(err)
		}

	case telegram.InsertCorrectPortButtonData:
		if err := c.service.CorrectPortCallback(chatID, messageID, text); err != nil {
			panic(err)
		}

	case telegram.InsertCorrectMailboxButtonData:
		if err := c.service.CorrectMailboxCallback(chatID, messageID, text); err != nil {
			panic(err)
		}

	case telegram.InsertCorrectUsernameButtonData:
		if err := c.service.CorrectUsernameCallback(chatID, messageID, text); err != nil {
			panic(err)
		}

	case telegram.InsertCorrectPasswordButtonData:
		if err := c.service.CorrectPasswordCallback(chatID, messageID, text); err != nil {
			panic(err)
		}

	case telegram.InsertSubmitButtonData:
		if err := c.service.SubmitCallback(chatID, messageID, text); err != nil {
			panic(err)
		}
	}
}
