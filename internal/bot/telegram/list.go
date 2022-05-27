package telegram

import (
	"fmt"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	listMsg = "Отслеживаемые ящики:"

	ListMailboxesPathPrefix = "/list"

	strMailboxFormat = "%s:%d %s %s"
)

func (t *TelegramService) SendMailboxesList(chatID int64, mailboxes []*pb.MailboxV1) error {

	msg := tgbotapi.NewMessage(chatID, listMsg)
	msg.ReplyMarkup = assembleKeyboard(mailboxes)

	_, err := t.bot.Send(msg)
	return err
}

func assembleKeyboard(mailboxes []*pb.MailboxV1) tgbotapi.InlineKeyboardMarkup {
	var rows [][]tgbotapi.InlineKeyboardButton

	for _, mailbox := range mailboxes {
		rows = append(rows,
			tgbotapi.NewInlineKeyboardRow(
				tgbotapi.NewInlineKeyboardButtonData(
					stringifyMailbox(mailbox),
					ListMailboxesPathPrefix+fmt.Sprintf("/%d", mailbox.MailboxID)+"/show",
				)))
	}
	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}

func stringifyMailbox(mailbox *pb.MailboxV1) string {
	return fmt.Sprintf(strMailboxFormat,
		mailbox.Host,
		mailbox.Port,
		mailbox.Username,
		mailbox.Mailbox,
	)
}
