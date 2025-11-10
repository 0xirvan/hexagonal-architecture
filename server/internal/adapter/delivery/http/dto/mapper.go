package dto

import (
	"errors"
	"net/http"

	"github.com/0xirvan/tdl-svelte-go/server/internal/core/domain"
)

// MapDomainError maps domain errors to HTTP status codes and messages(pure error string)
func MapDomainError(err error) (int, string) {
	switch {
	case errors.Is(err, domain.ErrTodoNotFound):
		return http.StatusNotFound, err.Error()
	case errors.Is(err, domain.ErrInvalidID):
		return http.StatusBadRequest, err.Error()
	case errors.Is(err, domain.ErrInvalidRequestBody):
		return http.StatusBadRequest, err.Error()
	}

	return http.StatusInternalServerError, err.Error()
}
