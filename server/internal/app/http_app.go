package app

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/0xirvan/tdl-svelte-go/server/internal/adapter/config"
	httpdelivery "github.com/0xirvan/tdl-svelte-go/server/internal/adapter/delivery/http"
)

type HTTPApp struct {
	router *httpdelivery.Router
	cfg    *config.AppContainer
}

func NewHTTPApp(
	router *httpdelivery.Router,
	cfg *config.AppContainer,
) *HTTPApp {
	return &HTTPApp{router: router, cfg: cfg}
}

func (a *HTTPApp) Run(ctx context.Context) error {
	addr := fmt.Sprintf("%s:%s", a.cfg.HTTP.Host, a.cfg.HTTP.Port)
	slog.Info("HTTP server starting", slog.String("addr", addr))

	errCh := make(chan error, 1)
	go func() {
		errCh <- a.router.Start(addr)
	}()

	// Wait for shutdown signal or server error
	select {
	case err := <-errCh:
		return fmt.Errorf("server stopped unexpectedly: %w", err)
	case <-waitForShutdownSignal():
		slog.Info("shutdown signal received, shutting down gracefully...")
	}

	// Graceful shutdown
	shutdownCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	if err := a.router.Shutdown(shutdownCtx); err != nil {
		return fmt.Errorf("error during graceful shutdown: %w", err)
	}

	slog.Info("server gracefully stopped")
	return nil
}

func waitForShutdownSignal() <-chan os.Signal {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	return sigCh
}
