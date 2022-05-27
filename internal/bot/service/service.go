package service

import (
	"github.com/aveplen/silicon-funnel/internal/bot/builder"
	"github.com/aveplen/silicon-funnel/internal/bot/client"
	"github.com/aveplen/silicon-funnel/internal/bot/telegram"
)

var _ TelegramCapabilities = (*telegram.TelegramService)(nil)

type TelegramCapabilities interface {
	TelegramAcquireCapabilities
	TelegramDetailCallbackCapabilities
	TelegramHelpCapabilities
	TelegramInsertCallbackCapabilities
	TelegramInsertCapabilities
	TelegramListCallbackCapabilities
	TelegramListCapabilities
	TelegramPollCapabilities
	TelegramStartCapabilities
}

var _ MailboxBuilderServiceCapabilities = (*builder.MailboxBuilderService)(nil)

type MailboxBuilderServiceCapabilities interface {
	MailboxBuilderServiceInsertCapabilities
	MailboxBuilderServiceInsertCallbackCapabilities
}

type Service struct {
	telegramService       TelegramCapabilities
	mailboxBuilderService MailboxBuilderServiceCapabilities
	client                client.ImapConcentratorClient
}

func NewService(
	telegramService TelegramCapabilities,
	mailboxBuilderService MailboxBuilderServiceCapabilities,
	client client.ImapConcentratorClient,
) *Service {

	return &Service{
		telegramService:       telegramService,
		mailboxBuilderService: mailboxBuilderService,
		client:                client,
	}
}
