package messages

import (
	"strconv"

	"github.com/aveplen/silicon-funnel/internal/bot/builder"
	"github.com/aveplen/silicon-funnel/internal/bot/service"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type MessageHandler struct {
	service *service.Service
}

func NewMessageHandler(service *service.Service) *MessageHandler {
	return &MessageHandler{
		service: service,
	}
}

func (m *MessageHandler) HandleMessage(update tgbotapi.Update) {
	chatID := update.Message.Chat.ID
	if m.service.BuilderMounted(chatID) {
		stage, err := m.service.GetBuilderStage(chatID)
		if err != nil {
			panic(err)
		}

		text := update.Message.Text

		switch stage {
		case builder.Host:
			if err := m.service.Host(chatID, text); err != nil {
				panic(err)
			}

		case builder.Port:
			port, _ := strconv.ParseInt(text, 10, 32)
			if err := m.service.Port(chatID, int32(port)); err != nil {
				panic(err)
			}

		case builder.Mailbox:
			if err := m.service.Mailbox(chatID, text); err != nil {
				panic(err)
			}

		case builder.Username:
			if err := m.service.Username(chatID, text); err != nil {
				panic(err)
			}

		case builder.Password:
			if err := m.service.Password(chatID, text); err != nil {
				panic(err)
			}

		case builder.CorrectHost:
			if err := m.service.CorrectHost(chatID, text); err != nil {
				panic(err)
			}

		case builder.CorrectPort:
			port, _ := strconv.ParseInt(text, 10, 32)
			if err := m.service.CorrectPort(chatID, int32(port)); err != nil {
				panic(err)
			}

		case builder.CorrectMailbox:
			if err := m.service.CorrectMailbox(chatID, text); err != nil {
				panic(err)
			}

		case builder.CorrectUsername:
			if err := m.service.CorrectUsername(chatID, text); err != nil {
				panic(err)
			}

		case builder.CorrectPassword:
			if err := m.service.CorrectPassword(chatID, text); err != nil {
				panic(err)
			}

		case builder.Submit:
			if err := m.service.Submit(chatID); err != nil {
				panic(err)
			}
		}
	}
}
