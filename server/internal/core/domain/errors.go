package domain

import "errors"

// global domain errors
var (
	ErrInvalidID          = errors.New("invalid id format")
	ErrInvalidRequestBody = errors.New("invalid request body")
)

type ValidationErrors map[string]string

func (v ValidationErrors) Error() string {
	return "validation failed"
}

// Todo domain errors
var (
	ErrTodoNotFound = errors.New("todo not found")
)
