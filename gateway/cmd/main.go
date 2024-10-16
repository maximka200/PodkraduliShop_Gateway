package main

import (
	"context"
	"fmt"
	"gateway/internal/client/product"
	"gateway/internal/config"
	"gateway/internal/server/handlers"
	"gateway/internal/server/server"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// init config
	cfg := config.MustReadConfig()
	// set secret in env
	config.SetEnvSecret(cfg.SecretKey)
	// init logger
	log := initLogger(cfg.Env)
	log.Info("logger and config successfully init")
	// run server
	serv := new(server.Server)
	log.Info(fmt.Sprintf("server run, port: %s", cfg.Port))
	// init context
	ctx, cancel := context.WithTimeout(context.Background(), cfg.Product.Timeout)
	defer cancel()

	// init grpc product client
	// todo debug: cannot read addres from config, ctx
	grpcProduct, err := product.NewClientProduct(log, cfg.Product.Addr, cfg.Product.Timeout, cfg.Product.RetryCount)
	if err != nil {
		log.Error(fmt.Sprintf("cannot run grpc client: %s", err))
		panic("cannot create grpc client")
	}

	// todo debug: cannot read addres from config
	grpcAuth, err := product.NewClientAuth(log, cfg.Auth.Addr, cfg.Auth.Timeout, cfg.Auth.RetryCount)
	if err != nil {
		log.Error(fmt.Sprintf("cannot run grpc client: %s", err))
		panic("cannot create grpc client")
	}

	grpcMethods := product.NewGRPCMethods(*grpcAuth, *grpcProduct)
	var grpcClient handlers.ClientMethods = grpcMethods
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
