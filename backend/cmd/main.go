package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/timsofteng/xyz-home-task/adapters/googleBooks"
	"github.com/timsofteng/xyz-home-task/adapters/httpHandlers"
	openlibrary "github.com/timsofteng/xyz-home-task/adapters/openLibrary"
	"github.com/timsofteng/xyz-home-task/config"
	"github.com/timsofteng/xyz-home-task/internal/httpServer"
	"github.com/timsofteng/xyz-home-task/internal/logger"
	"github.com/timsofteng/xyz-home-task/service"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		slog.Error("top level", "details", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.ReadConfig()
	if err != nil {
		return fmt.Errorf("failed to read config: %w", err)
	}

	logger := logger.New(cfg.LogLevel)

	openLibClient := openlibrary.New(logger)

	googleBooksClient := googleBooks.New(logger, openLibClient)

	service := service.New(googleBooksClient)
	httpHandlers := httpHandlers.New(logger, service)

	ctx, cancel := signal.NotifyContext(
		context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()
	g, gCtx := errgroup.WithContext(ctx)

	server, err := httpServer.New(
		ctx,
		logger, httpHandlers,
		httpServer.Cfg{Host: cfg.HTTPServerHost, Port: cfg.HTTPServerPort},
	)
	if err != nil {
		return fmt.Errorf("failed to create http server: %w", err)
	}

	g.Go(server.Start)
	logger.Info(
		"http server has been started",
		"host", cfg.HTTPServerHost,
		"port", cfg.HTTPServerPort,
	)

	g.Go(func() error {
		<-gCtx.Done()

		shutdownCtx, shutdownCancel := context.WithTimeout(
			context.Background(), 5*time.Second)
		defer shutdownCancel()

		err := server.Stop(shutdownCtx)
		if err != nil {
			return fmt.Errorf("failed to gracefully stop the server: %w", err)
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return fmt.Errorf("failed to shutdown gracefully: %v", err)
	}

	return nil
}
