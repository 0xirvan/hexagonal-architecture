package http

import (
	"strconv"

	"github.com/labstack/echo/v4"
)

// WriteErrorResponse writes a standardized error response to the client
func WriteErrorResponse(c echo.Context, status int, err error) error {
	return c.JSON(status, map[string]any{"error": err.Error()})
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
