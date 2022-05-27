package callback

import (
	"strconv"
	"strings"

	"github.com/aveplen/silicon-funnel/internal/bot/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var _ ServiceListCallbackCapabilities = (*service.Service)(nil)

type ServiceListCallbackCapabilities interface {
	MailboxDetail(chatID int64, mailboxID int64, messageID int) error
	DeleteMailbox(chatID int64, mailboxID int64, messageID int) error
}

func (i *CallbackQueryHandler) List(update tgbotapi.Update) {
	chatID := update.CallbackQuery.Message.Chat.ID
	messageID := update.CallbackQuery.Message.MessageID

	pathParts := strings.Split(update.CallbackData(), "/")
	strMailboxID := pathParts[2]

	mailboxID, err := strconv.ParseInt(strMailboxID, 10, 64)
	if err != nil {
		panic(err)
	}

	if err := i.service.MailboxDetail(chatID, mailboxID, messageID); err != nil {
		panic(err)
	}
}
