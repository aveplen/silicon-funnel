package server

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aveplen/silicon-funnel/internal/bot/builder"
	"github.com/aveplen/silicon-funnel/internal/bot/callback"
	"github.com/aveplen/silicon-funnel/internal/bot/client"
	"github.com/aveplen/silicon-funnel/internal/bot/commands"
	"github.com/aveplen/silicon-funnel/internal/bot/config"
	"github.com/aveplen/silicon-funnel/internal/bot/messages"
	"github.com/aveplen/silicon-funnel/internal/bot/middleware"
	"github.com/aveplen/silicon-funnel/internal/bot/router"
	"github.com/aveplen/silicon-funnel/internal/bot/service"
	"github.com/aveplen/silicon-funnel/internal/bot/telegram"
	"github.com/aveplen/silicon-funnel/internal/server/errl"
	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func Start(cfg *config.Config) {
	ctx, ctxCancel := context.WithCancel(context.Background())

	// ================================= bot ==================================
	bot, err := tgbotapi.NewBotAPI(cfg.TgBot.Token)
	if err != nil {
		logrus.Fatal(err)
	}

	notifications := make(chan *pb.NotificationV1, 100)
	errs := make(chan error, 100)
	telegramService := telegram.NewTelegramService(bot, notifications, errs)

	// ============================== grpc client =============================
	concentratorClient, err := client.NewImapConcentratorClient(cfg, notifications, errs)
	if err != nil {
		logrus.Fatal(err)
	}

	// =============================== builders ===============================
	mailboxBuilderService := builder.NewMailboxBuilderService()

	service := service.NewService(telegramService, mailboxBuilderService, concentratorClient)

	// =============================== handlers ===============================
	c := commands.NewCommandHandler(service)
	m := messages.NewMessageHandler(service)
	cb := callback.NewCallbackQueryHandler(service)

	// =============================== router =================================
	r := router.NewUpdateRouter(bot)
	r.Use(middleware.DefaultPanicRecoveryMiddleware)

	// ============================== commands ===============================
	for k, v := range map[string]router.UpdateHandlerFunc{
		"start":   c.Start,
		"insert":  c.Insert,
		"help":    c.Help,
		"list":    c.List,
		"poll":    c.Poll,
		"acquire": c.Acquire,
	} {
		r.RouteCommandFunc(k, v)
	}

	// ============================== callbacks ===============================
	for k, v := range map[string]router.UpdateHandlerFunc{
		"/insert": cb.Insert,
		"/list":   cb.List,
		"/detail": cb.Detail,
	} {
		r.RouteCallbackQuery(k, v)
	}

	// =============================== messages ===============================
	r.RouteFunc(func(update tgbotapi.Update) bool { return true }, m.HandleMessage)

	// ============================== listeners ===============================
	var group errgroup.Group

	group.Go(func() error {
		return concentratorClient.Start(ctx)
	})

	group.Go(func() error {
		return telegramService.Start(ctx)
	})

	group.Go(func() error {
		return r.Run(ctx, tgbotapi.UpdateConfig{})
	})

	group.Go(func() error {
		return http.ListenAndServe(":9090", LivenessProbe(http.NewServeMux()))
	})

	group.Go(func() error {
		return errl.ListenForErrors(ctx, errs)
	})

	// =========================== graceful shutdown ==========================
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	errch := make(chan error)
	go func() {
		errch <- group.Wait()
	}()

	select {
	case <-sigc:
		ctxCancel()
		time.Sleep(5 * time.Second)

		logrus.Info("exited cleanly")

	case err := <-errch:
		ctxCancel()
		time.Sleep(5 * time.Second)

		logrus.Fatal(err)
	}
}
