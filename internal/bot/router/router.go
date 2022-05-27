//go:generate mkdir -p mock
//go:generate minimock -o ./mock/ -s .go -g

package router

import (
	"context"
	"fmt"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IRouter interface {
	RouteMessagesFunc(handlers ...UpdateHandlerFunc)
}

type UpdateFilter func(tgbotapi.Update) bool

type UpdateHandlerFunc func(tgbotapi.Update)

type Route struct {
	filter   UpdateFilter
	handlers []UpdateHandlerFunc
}

type Middleware func(UpdateHandlerFunc) UpdateHandlerFunc

var _ UpdateRouter = (*UpdateRouterImpl)(nil)

type UpdateRouter interface {
	Run(ctx context.Context, config tgbotapi.UpdateConfig) error
	Use(middleware Middleware)
	RouteFunc(filter UpdateFilter, handlers ...UpdateHandlerFunc)
	RouteMessagesFunc(handlers ...UpdateHandlerFunc)
	RouteCommandFunc(command string, handlers ...UpdateHandlerFunc)
	RouteCallbackQuery(path string, handlers ...UpdateHandlerFunc)
}

type UpdateRouterImpl struct {
	bot        *tgbotapi.BotAPI
	routes     []Route
	middleware []Middleware
}

func NewUpdateRouter(bot *tgbotapi.BotAPI) *UpdateRouterImpl {
	return &UpdateRouterImpl{
		bot:    bot,
		routes: []Route{},
	}
}

func (r *UpdateRouterImpl) Run(ctx context.Context, config tgbotapi.UpdateConfig) error {
	updates := r.bot.GetUpdatesChan(config)

	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("context done")

		case update := <-updates:
			r.handle(update)
		}
	}
}

func (r *UpdateRouterImpl) Use(middleware Middleware) {
	r.middleware = append(r.middleware, middleware)
}

func (r *UpdateRouterImpl) RouteFunc(filter UpdateFilter, handlers ...UpdateHandlerFunc) {
	r.routes = append(r.routes, Route{filter: filter, handlers: handlers})
}

func (r *UpdateRouterImpl) handle(update tgbotapi.Update) {
	for _, route := range r.routes {
		if route.filter(update) {
			for _, handlerFunc := range route.handlers {
				for i := range r.middleware {
					handlerFunc = r.middleware[len(r.middleware)-i-1](handlerFunc)
				}
				handlerFunc(update)
				return
			}
		}
	}
}

func (r *UpdateRouterImpl) RouteMessagesFunc(handlers ...UpdateHandlerFunc) {
	r.RouteFunc(func(update tgbotapi.Update) bool {
		return update.Message != nil
	}, handlers...)
}

func (r *UpdateRouterImpl) RouteCommandFunc(command string, handlers ...UpdateHandlerFunc) {
	if len(command) == 0 {
		panic("command Route should not be empty string")
	}

	parts := strings.Split(command, " ")
	if len(parts) != 1 {
		panic("command Route should only contain one word")
	}

	r.RouteFunc(func(update tgbotapi.Update) bool {
		return update.Message != nil &&
			update.Message.IsCommand() &&
			update.Message.Command() == parts[0]
	}, handlers...)
}

func (r *UpdateRouterImpl) RouteCallbackQuery(path string, handlers ...UpdateHandlerFunc) {
	r.RouteFunc(func(update tgbotapi.Update) bool {
		return update.CallbackQuery != nil &&
			update.CallbackQuery.Data[:len(path)] == path
	}, handlers...)
}
