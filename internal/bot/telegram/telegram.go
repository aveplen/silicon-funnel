//go:generate mkdir -p mock
//go:generate minimock -o ./mock/ -s .go -g

package telegram

import (
	"context"
	"fmt"

	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramService struct {
	bot           *tgbotapi.BotAPI
	notifications <-chan *pb.NotificationV1
	errors        chan error
}

func NewTelegramService(
	bot *tgbotapi.BotAPI,
	notifications <-chan *pb.NotificationV1,
	errors chan error,
) *TelegramService {

	return &TelegramService{
		bot:           bot,
		notifications: notifications,
		errors:        errors,
	}
}

func (t *TelegramService) Start(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context done")
		case notification := <-t.notifications:
			if err := t.SendText(notification.ChatID, notification.Title+"\n"+notification.Body); err != nil {
				return err
			}
		}
	}
}

func (t *TelegramService) SendText(chatID int64, text string) error {
	msg := tgbotapi.NewMessage(chatID, text)

	if _, err := t.bot.Send(msg); err != nil {
		return err
	}

	return nil
}
