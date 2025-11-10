package helper

import (
	"strconv"

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

// StrToUint converts a string to uint
func StrToUint(s string) (uint, error) {
	val, err := strconv.Atoi(s)
	return uint(val), err
}

// StrToInt converts a string to int
func StrToInt(s string) (int, error) {
	return strconv.Atoi(s)
}
