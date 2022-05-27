package callback

import "github.com/aveplen/silicon-funnel/internal/bot/service"

var _ ServiceCapabilities = (*service.Service)(nil)

type ServiceCapabilities interface {
	ServiceDetailCallbackCapabilities
	ServiceInsertCallbackCapabilities
	ServiceListCallbackCapabilities
}

type CallbackQueryHandler struct {
	service ServiceCapabilities
}

func NewCallbackQueryHandler(service ServiceCapabilities) *CallbackQueryHandler {
	return &CallbackQueryHandler{
		service: service,
	}
}
