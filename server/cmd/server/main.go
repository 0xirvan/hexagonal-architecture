package main

import (
	"log/slog"

	"github.com/0xirvan/tta-svelte-go/server/internal/adapter/config"
)

func main() {
	if err := run(); err != nil {
		slog.Error("application terminated with error", slog.String("error", err.Error()))
	}
}

func run() error {
	// Load config
	cfg, err := config.New()
	if err != nil {
		return err
	}

	slog.Info("config loaded successfully", slog.Any("config", cfg))
	return nil
}
