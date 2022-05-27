package telegram

func (t *TelegramService) NotifyAboutStart(chatID int64) error {
	return t.SendHelpText(chatID)
}
