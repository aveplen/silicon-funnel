package telegram

import (
	"fmt"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	detailPathPrefix = "detail"
	detailPathFormat = "/%s/%d/%s"

	detailReturnButtonText = "Назад"
	detailDeleteButtonText = "Удалить"

	DetailReturnButtonData = "return"
	DetailDeleteButtonData = "delete"
)

func (t *TelegramService) ChangeKeyboardToDetail(chatID int64, messageID int, mailbox *pb.MailboxV1) error {
	text := stringifyMailbox(mailbox)

	detailInlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
		// return button
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				detailReturnButtonText,
				fmt.Sprintf(detailPathFormat,
					detailPathPrefix,
					mailbox.MailboxID,
					DetailReturnButtonData))),

		// delete button
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				detailDeleteButtonText,
				fmt.Sprintf(detailPathFormat,
					detailPathPrefix,
					mailbox.MailboxID,
					DetailDeleteButtonData))),
	)

	msg := tgbotapi.NewEditMessageTextAndMarkup(chatID, messageID, text, detailInlineKeyboard)

	if _, err := t.bot.Send(msg); err != nil {
		return err
	}
	return nil
}

func (t *TelegramService) ReturnToMailboxes(chatID int64, messageID int, mailboxes []*pb.MailboxV1) error {
	msg := tgbotapi.NewEditMessageTextAndMarkup(chatID, messageID, listMsg, assembleKeyboard(mailboxes))

	if _, err := t.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
