package telegram

import (
	"fmt"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	insertFinalMessage = "Ящик добавлен"

	insertCorrectPathPrefix = "insert"
	insertPathFormat        = "/%s/%s"

	insertSubmitButtonText          = "Подтвердить"
	insertCorrectHostButtonText     = "Изменить host"
	insertCorrectPortButtonText     = "Изменить port"
	insertCorrectMailboxButtonText  = "Изменить mailbox"
	insertCorrectUsernameButtonText = "Изменить username"
	insertCorrectPasswordButtonText = "Изменить password"

	insertHostQuestion     = "Host? (пример: imap.yandex.ru)"
	insertPortQuestion     = "Port? (пример: 993)"
	insertMailboxQuestion  = "Mailbox? (пример: INBOX)"
	insertUsernameQuestion = "Username? (пример: imap_username)"
	insertPasswordQuestion = "Password? (пример: imap_password)"

	InsertSubmitButtonData          = "submit"
	InsertCorrectHostButtonData     = "host"
	InsertCorrectPortButtonData     = "port"
	InsertCorrectMailboxButtonData  = "mailbox"
	InsertCorrectUsernameButtonData = "username"
	InsertCorrectPasswordButtonData = "password"
)

var (
	insertInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
		// change host button
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				insertCorrectHostButtonText,
				fmt.Sprintf(insertPathFormat,
					insertCorrectPathPrefix,
					InsertCorrectHostButtonData))),

		// change port button
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				insertCorrectPortButtonText,
				fmt.Sprintf(insertPathFormat,
					insertCorrectPathPrefix,
					InsertCorrectPortButtonData))),

		// change mailbox button
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				insertCorrectMailboxButtonText,
				fmt.Sprintf(insertPathFormat,
					insertCorrectPathPrefix,
					InsertCorrectMailboxButtonData))),

		// change username button
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				insertCorrectUsernameButtonText,
				fmt.Sprintf(insertPathFormat,
					insertCorrectPathPrefix,
					InsertCorrectUsernameButtonData))),

		// change password button
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				insertCorrectPasswordButtonText,
				fmt.Sprintf(insertPathFormat,
					insertCorrectPathPrefix,
					InsertCorrectPasswordButtonData))),

		// change submit button
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(
				insertSubmitButtonText,
				fmt.Sprintf(insertPathFormat,
					insertCorrectPathPrefix,
					InsertSubmitButtonData))),
	)
)

func (t *TelegramService) AskForHost(chatID int64) error {
	return t.SendText(chatID, insertHostQuestion)
}

func (t *TelegramService) AskForPort(chatID int64) error {
	return t.SendText(chatID, insertPortQuestion)
}

func (t *TelegramService) AskForMailbox(chatID int64) error {
	return t.SendText(chatID, insertMailboxQuestion)
}

func (t *TelegramService) AskForUsername(chatID int64) error {
	return t.SendText(chatID, insertUsernameQuestion)
}

func (t *TelegramService) AskForPassword(chatID int64) error {
	return t.SendText(chatID, insertPasswordQuestion)
}

func (t *TelegramService) AskForSubmit(chatID int64, request *pb.MailboxV1) error {
	return t.sendTextWithKeyboard(chatID, fmt.Sprintf("%v", request), insertInlineKeyboard)
}

func (t *TelegramService) InsertSubmitResponse(chatID int64, messageID int) error {
	if err := t.SendText(chatID, insertFinalMessage); err != nil {
		return err
	}
	return t.CloseKeyboard(chatID, messageID, insertFinalMessage)
}
