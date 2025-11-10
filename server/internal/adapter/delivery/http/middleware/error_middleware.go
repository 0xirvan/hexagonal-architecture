package middleware

import (
	"net/http"

	"github.com/0xirvan/hexagonal-architecture/server/internal/adapter/delivery/http/dto"
	"github.com/0xirvan/hexagonal-architecture/server/internal/adapter/delivery/http/helper"
	"github.com/0xirvan/hexagonal-architecture/server/internal/core/domain"
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

		// validation errors
		if vErr, ok := err.(domain.ValidationErrors); ok {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":  "validation_error",
				"fields": vErr,
			})
		}

		// domains / business errors
		status, title := dto.MapDomainError(err)
		return helper.WriteErrorResponse(c, status, title, err.Error())
	}
}
