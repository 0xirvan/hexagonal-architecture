package helper

import (
	"strconv"

	"github.com/0xirvan/hexagonal-architecture/server/internal/adapter/delivery/http/dto"
	"github.com/labstack/echo/v4"
)

// WriteErrorResponse writes a standardized error response to the client
// with the given status code, title, and detail message
func WriteErrorResponse(c echo.Context, status int, title string, detail string) error {
	return c.JSON(status, map[string]any{
		"error": map[string]any{
			"title":  title,
			"detail": detail,
			"status": status,
		},
	})
}

// WriteList writes a standardized list response to the client
func WriteList[T any](c echo.Context, status int, items []T) error {
	return c.JSON(status, map[string]any{
		"data": items,
	})
}

// WriteSingle writes a standardized single item response to the client
func WriteSingle[T any](c echo.Context, status int, item T) error {
	return c.JSON(status, dto.ToSingleResponse(item))
}

// WritePaginated writes a standardized paginated response to the client
func WritePaginated[T any](c echo.Context, status int, items []T, total, page, size int) error {
	return c.JSON(status, dto.ToPaginatedResponse(items, total, page, size))
}

// StrToUint converts a string to uint
func StrToUint(s string) (uint, error) {
	val, err := strconv.Atoi(s)
	return uint(val), err
}

// StrToInt converts a string to int
func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
