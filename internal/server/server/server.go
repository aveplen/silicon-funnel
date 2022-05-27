package server

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/aveplen/silicon-funnel/internal/server/client"
	"github.com/aveplen/silicon-funnel/internal/server/config"
	"github.com/aveplen/silicon-funnel/internal/server/errl"
	"github.com/aveplen/silicon-funnel/internal/server/poller"
	"github.com/aveplen/silicon-funnel/internal/server/repo"
	"github.com/aveplen/silicon-funnel/internal/server/service"
	pb "github.com/aveplen/silicon-funnel/pkg/imap_concentrator/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/sirupsen/logrus"
	"github.com/tmc/grpc-websocket-proxy/wsproxy"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Start(cfg *config.Config) {
	ctx, ctxCancel := context.WithCancel(context.Background())
	errs := make(chan error, 100)

	// ============================== repository ==============================
	repository, err := repo.NewRepository(ctx, cfg, errs)
	if err != nil {
		log.Fatal(err)
	}

	// ================================ client ================================
	notifications := make(chan *pb.NotificationV1, 100)

	imapClient, err := client.NewImapClient(repository, notifications, errs)
	if err != nil {
		log.Fatal(err)
	}
	pollerService, err := poller.NewPoller(cfg, imapClient)
	if err != nil {
		log.Fatal(err)
	}

	go pollerService.Start(ctx)

	// ============================= grpc server ==============================
	listen, err := net.Listen("tcp", cfg.Imapc.Addr)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	server, err := service.NewImapConcentratorServer(*cfg.Imapc.Whitelist, repository, imapClient, notifications)
	if err != nil {
		log.Fatal(err)
	}

	pb.RegisterImapConcentratorServer(grpcServer, server)

	// =============================== gateway ================================
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	// ============================== listeners ===============================
	var group errgroup.Group

	// grpc subscribers
	group.Go(func() error {
		logrus.Infof("grpc server listening on %s", cfg.Imapc.Addr)
		return server.ListenForSubscribers(ctx)
	})

	// repository offset bumps
	group.Go(func() error {
		return repository.ListenForBumps(ctx)
	})

	// grpc gateway control
	group.Go(func() error {
		return pb.RegisterImapConcentratorHandlerFromEndpoint(ctx, mux, cfg.Imapc.Addr, opts)
	})

	// grpc gateway websocket stream
	group.Go(func() error {
		logrus.Infof("gateway websocket listening on 0.0.0.0:8080")
		return http.ListenAndServe("0.0.0.0:8080", LivenessProbe(wsproxy.WebsocketProxy(mux)))
	})

	// error logger
	group.Go(func() error {
		return errl.ListenForErrors(ctx, errs)
	})

	// grpc server
	group.Go(func() error {
		errch := make(chan error)
		go func() {
			errch <- grpcServer.Serve(listen)
		}()

		select {
		case <-ctx.Done():
			grpcServer.Stop()
			return fmt.Errorf("context done")
		case err := <-errch:
			return err
		}
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
