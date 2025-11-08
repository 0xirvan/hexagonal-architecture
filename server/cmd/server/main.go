package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0xirvan/tta-svelte-go/server/internal/adapter/config"
	"github.com/0xirvan/tta-svelte-go/server/internal/adapter/delivery/http"
)

func main() {
	if err := run(); err != nil {
		slog.Error("application terminated with error", slog.String("error", err.Error()))
		os.Exit(1)
	}
}

func run() error {
	// Load config
	cfg, err := config.New()
	if err != nil {
		return err
	}

	slog.Info("config loaded successfully", slog.Any("config", cfg))

	// Init router
	router := http.NewRouter(cfg.HTTP)

	// Start server in antoher goroutine
	listenAddr := fmt.Sprintf("%s:%s", cfg.HTTP.Host, cfg.HTTP.Port)
	slog.Info("starting HTTP server on", slog.String("address", listenAddr))

	serverErrCh := make(chan error, 1)
	go func() {
		serverErrCh <- router.Start(listenAddr)
	}()

	// Graceful shutdown
	// Wait for server error or termination signal
	select {
	case <-serverErrCh:
		return fmt.Errorf("HTTP server stopped unexpectedly: %w", <-serverErrCh)
	case <-waitForSignal():
		slog.Info("termination signal received, shutting down server...")
	}

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := router.Shutdown(shutdownCtx); err != nil {
		return err
	}
	return nil
}

// waitForSignal waits for OS interrupt or termination signals
func waitForSignal() <-chan os.Signal {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	return sigCh
}
