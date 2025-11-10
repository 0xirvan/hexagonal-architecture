package todo

import (
	"context"

	"github.com/0xirvan/hexagonal-architecture/server/internal/core/domain"
	"github.com/0xirvan/hexagonal-architecture/server/internal/core/port"
)

type ListTodosUsecase struct {
	Repo port.TodoRepository
}

// Execute retrieves all todo items
func (uc *ListTodosUsecase) Execute(ctx context.Context) ([]*domain.Todo, error) {
	return uc.Repo.FindAll(ctx)
}
