//go:generate mkdir -p mock
//go:generate minimock -o ./mock/ -s .go -g

package commands

type ServiceCapabilities interface {
	ServiceAcquireCapabilities
	ServiceStartCapabilities
	ServiceHelpCapabilities
	ServiceInsertCapabilities
	ServiceInsertCapabilities
	ServiceListCapabilities
	ServicePollCapabilities
}

type CommandHandler struct {
	service ServiceCapabilities
}

func NewCommandHandler(service ServiceCapabilities) *CommandHandler {
	return &CommandHandler{
		service: service,
	}
}
