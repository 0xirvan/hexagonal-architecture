package middleware

import (
	"github.com/0xirvan/tdl-svelte-go/server/internal/adapter/delivery/http/dto"
	"github.com/0xirvan/tdl-svelte-go/server/internal/adapter/delivery/http/helper"
	"github.com/labstack/echo/v4"
)

// ErrorMiddleware is an Echo middleware that handles errors returned by handlers
// and maps them to standardized HTTP error responses
func ErrorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := next(c)
		if err == nil {
			return nil
		}

		status, title := dto.MapDomainError(err)
		return helper.WriteErrorResponse(c, status, title, err.Error())
	}
}
