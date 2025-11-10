package main

import (
	"context"
	"log/slog"
	"os"

	"github.com/0xirvan/tdl-svelte-go/server/internal/adapter/config"
	"github.com/0xirvan/tdl-svelte-go/server/internal/app"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("failed to load config", slog.String("error", err.Error()))
		os.Exit(1)
	}

	httpApp, err := app.InitializeHTTPApp(context.Background(), cfg)
	if err != nil {
		slog.Error("failed to build app", slog.String("error", err.Error()))
		os.Exit(1)
	}

	if err := httpApp.Run(context.Background()); err != nil {
		slog.Error("server exited with error", slog.String("error", err.Error()))
		os.Exit(1)
	}
}
