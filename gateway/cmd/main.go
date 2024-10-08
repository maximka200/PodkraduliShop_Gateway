package main

import (
	"context"
	"fmt"
	"geteway/internal/client/product"
	"geteway/internal/config"
	"geteway/internal/server/handlers"
	"geteway/internal/server/server"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init config
	cfg := config.MustReadConfig()
	// init logger
	log := initLogger(cfg.Env)
	log.Info("logger and config successfully init")
	// run server
	serv := new(server.Server)
	log.Info(fmt.Sprintf("server run, port: %s", cfg.Port))
	// init context
	ctx, cancel := context.WithTimeout(context.Background(), cfg.ProductGRPC.Timeout)
	defer cancel()
	// init grpc product client
	grpc, err := product.NewClientProduct(log, cfg.ProductGRPC.Addr, cfg.ProductGRPC.Timeout, cfg.ProductGRPC.RetryCount)
	if err != nil {
		log.Error(fmt.Sprintf("cannot run grpc client: %s", err))
		panic("cannot create grpc client")
	}
	var grpcClient handlers.ClientMethods = grpc
	// init handler
	handler := handlers.NewHandler(log, grpcClient, ctx)
	go func() {
		if err := serv.Run(cfg, handler.InitRouter()); err != nil {
			log.Error(fmt.Sprintf("cannot run server: %s", err))
			panic("cannot run server")
		}
	}()
	// shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	if err := serv.Shutdown(ctx); err != nil {
		log.Error(fmt.Sprintf("an error occurred while executing graceful shutdown: %s", err))
	}
}

func initLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case "local":
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)

	case "dev":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case "prod":
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
