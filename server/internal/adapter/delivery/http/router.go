package http

import (
	"net/http"
	"strings"

	"github.com/0xirvan/hexagonal-architecture/server/internal/adapter/config"
	"github.com/0xirvan/hexagonal-architecture/server/internal/adapter/delivery/http/handler"
	customMiddleware "github.com/0xirvan/hexagonal-architecture/server/internal/adapter/delivery/http/middleware"
	"github.com/0xirvan/hexagonal-architecture/server/internal/adapter/delivery/http/routes"
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
	todoHandler *handler.TodoHandler,
) *Router {
	e := echo.New()

	e.Validator = NewRequestValidator("en")

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(customMiddleware.ErrorMiddleware)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: strings.Split(cfg.AllowedOrigins, ","),
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete, http.MethodOptions},
	}))

	v1 := e.Group("/api/v1")

	routes.RegisterTodoRoutes(v1, todoHandler)

	return &Router{e}
}
