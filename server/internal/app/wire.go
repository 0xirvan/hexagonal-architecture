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

// Infrastructure providers
var infraSet = wire.NewSet(
	ProvideHTTPConfig,
)

// Repository layer
var repositorySet = wire.NewSet(
	inmemory.NewTodoRepository,
)

// Use case layer
var useCaseSet = wire.NewSet(
	todo.NewService,
)

// Delivery layer
var deliverySet = wire.NewSet(
	httpdelivery.NewTodoHandler,
	httpdelivery.NewRouter,
)

// All sets
var appSet = wire.NewSet(
	infraSet,
	repositorySet,
	useCaseSet,
	deliverySet,
	NewHTTPApp,
)

func InitializeHTTPApp(ctx context.Context, cfg *config.AppContainer) (*HTTPApp, error) {
	wire.Build(appSet)
	return &HTTPApp{}, nil
}
