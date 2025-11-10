package todo

import (
	"context"

	"github.com/0xirvan/hexagonal-architecture/server/internal/core/domain"
	"github.com/0xirvan/hexagonal-architecture/server/internal/core/port"
)

type MarkTodoUndoneUsecase struct {
	Repo port.TodoRepository
}

// Execute marks a todo item as not done by its ID, also clearing the CompletedAt timestamp
func (uc *MarkTodoUndoneUsecase) Execute(ctx context.Context, id uint) (*domain.Todo, error) {
	todo, err := uc.Repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	todo.IsDone = false
	todo.CompletedAt = nil

	return uc.Repo.Update(ctx, todo)
}
