package telegram

import "fmt"

const acquireMsgFormat = "Ваш идентификатор чата: %d"

func (t *TelegramService) ResponseToAcquire(chatID int64) error {
	return t.SendText(chatID, fmt.Sprintf(acquireMsgFormat, chatID))
}
