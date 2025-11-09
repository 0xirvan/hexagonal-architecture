package domain

import "errors"

// global domain errors
var (
	ErrInvalidID          = errors.New("invalid id format")
	ErrInvalidRequestBody = errors.New("invalid request body")
)

// Todo domain errors
var (
	ErrTodoNotFound = errors.New("todo not found")
)
