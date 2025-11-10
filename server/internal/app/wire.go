//go:build wireinject
// +build wireinject

package app

import (
	"context"

	"github.com/google/wire"

	"github.com/0xirvan/tdl-svelte-go/server/internal/adapter/config"
	httpdelivery "github.com/0xirvan/tdl-svelte-go/server/internal/adapter/delivery/http"
	"github.com/0xirvan/tdl-svelte-go/server/internal/adapter/persistence/inmemory"
	"github.com/0xirvan/tdl-svelte-go/server/internal/core/usecase/todo"
)

func ProvideHTTPConfig(cfg *config.AppContainer) *config.HTTP {
	return cfg.HTTP
}

var todoSet = wire.NewSet(
	inmemory.NewTodoRepository,
	todo.NewService,
	httpdelivery.NewTodoHandler,
)

func InitializeHTTPApp(ctx context.Context, cfg *config.AppContainer) (*HTTPApp, error) {
	wire.Build(
		todoSet,
		ProvideHTTPConfig,
		httpdelivery.NewRouter,
		NewHTTPApp,
	)
	return &HTTPApp{}, nil
}
