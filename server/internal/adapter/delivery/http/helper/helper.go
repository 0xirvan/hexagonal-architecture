package helper

import (
	"strconv"

	"github.com/0xirvan/tta-svelte-go/server/internal/adapter/delivery/http/dto"
	"github.com/labstack/echo/v4"
)

// WriteErrorResponse writes a standardized error response to the client
// with the given status code, title, and detail message
// ref https://tools.ietf.org/html/rfc7807
func WriteErrorResponse(c echo.Context, status int, title string, detail string) error {
	return c.JSON(status, dto.ErrorResponse{
		Type:     "about:blank", // todo: customize error type URLs as needed
		Title:    title,
		Status:   status,
		Detail:   detail,
		Instance: c.Path(),
	})
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
