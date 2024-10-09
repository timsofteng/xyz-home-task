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
	"github.com/timsofteng/xyz-home-task/adapters/openLibrary"
	"github.com/timsofteng/xyz-home-task/internal/env"
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
	host := env.MustGet("HOST")
	httpServerPort := env.MustGet("HTTP_SERVER_PORT")

	logLevel := os.Getenv("LOG_LEVEL")
	logger := logger.New(logLevel)

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
		httpServer.Cfg{Host: host, Port: httpServerPort},
	)

	if err != nil {
		return fmt.Errorf("failed to create http server: %w", err)
	}

	g.Go(server.Start)
	logger.Info(
		"http server has been started",
		"host", host,
		"port", httpServerPort,
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
