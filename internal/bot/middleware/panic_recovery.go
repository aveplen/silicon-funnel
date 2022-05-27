package middleware

import (
	"github.com/aveplen/silicon-funnel/internal/bot/router"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
)

func DefaultPanicRecoveryMiddleware(next router.UpdateHandlerFunc) router.UpdateHandlerFunc {
	return func(update tgbotapi.Update) {

		defer func() {
			if err := recover(); err != nil {
				logrus.Warn(err)
			}
		}()

		next(update)
	}
}
