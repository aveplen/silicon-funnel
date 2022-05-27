package callback

import (
	"strconv"
	"strings"

	"github.com/aveplen/silicon-funnel/internal/bot/service"
	"github.com/aveplen/silicon-funnel/internal/bot/telegram"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _ ServiceDetailCallbackCapabilities = (*service.Service)(nil)

type ServiceDetailCallbackCapabilities interface {
	ReturnToList(chatID int64, mailboxID int, messageID int) error
	DeleteMailbox(chatID int64, mailboxID int64, messageID int) error
}

func (n *CallbackQueryHandler) Detail(update tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	messageID := update.CallbackQuery.Message.MessageID

	pathParts := strings.Split(update.CallbackData(), "/")

	strMailboxID := pathParts[2]
	mailboxID, err := strconv.ParseInt(strMailboxID, 10, 64)
	if err != nil {
		panic(err)
	}

	action := pathParts[3]

	switch action {
	case telegram.DetailDeleteButtonData:
		if err := n.service.DeleteMailbox(chatID, mailboxID, messageID); err != nil {
			panic(err)
		}

	case telegram.DetailReturnButtonData:
		if err := n.service.ReturnToList(chatID, int(mailboxID), messageID); err != nil {
			panic(err)
		}
	}
}
