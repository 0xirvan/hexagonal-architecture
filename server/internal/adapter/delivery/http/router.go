package http

import (
	"net/http"
	"strings"

	"github.com/0xirvan/tdl-svelte-go/server/internal/adapter/config"
	customMiddleware "github.com/0xirvan/tdl-svelte-go/server/internal/adapter/delivery/http/middleware"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Router wraps the echo.Echo instance
type Router struct {
	*echo.Echo
}

// NewRouter creates a new Echo router instance
func NewRouter(
	cfg *config.HTTP,
) *Router {
	e := echo.New()

	setupMiddleware(e, cfg)

	return &Router{e}
}

// setupMiddleware configures the middleware for the Echo instance
func setupMiddleware(e *echo.Echo, cfg *config.HTTP) {
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(customMiddleware.ErrorMiddleware)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(cfg.AllowedOrigins, ","),
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
	}))
}
