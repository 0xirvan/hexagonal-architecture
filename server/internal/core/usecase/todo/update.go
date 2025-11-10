package todo

import (
	"context"

	"github.com/0xirvan/tdl-svelte-go/server/internal/core/domain"
	"github.com/0xirvan/tdl-svelte-go/server/internal/core/port"
)

type UpdateTodoUsecase struct {
	Repo port.TodoRepository
}

// Execute updates the title and/or description of a todo item by its ID
func (uc *UpdateTodoUsecase) Execute(
	ctx context.Context,
	id uint,
	title *string,
	desc *string,
) (*domain.Todo, error) {
	todo, err := uc.Repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if title != nil {
		todo.Title = *title
	}

	if desc != nil {
		todo.Description = *desc
	}

	return uc.Repo.Update(ctx, todo)
}
