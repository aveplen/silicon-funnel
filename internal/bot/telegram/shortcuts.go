package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (t *TelegramService) sendTextWithKeyboard(
	chatID int64,
	text string,
	markup tgbotapi.InlineKeyboardMarkup,
) error {

	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = markup

	_, err := t.bot.Send(msg)
	return err
}

func (t *TelegramService) CloseKeyboard(chatID int64, messageID int, text string) error {
	msg := tgbotapi.NewEditMessageTextAndMarkup(chatID, messageID, text, tgbotapi.InlineKeyboardMarkup{
		InlineKeyboard: make([][]tgbotapi.InlineKeyboardButton, 0),
	})
	if _, err := t.bot.Send(msg); err != nil {
		return err
	}
	return nil
}
