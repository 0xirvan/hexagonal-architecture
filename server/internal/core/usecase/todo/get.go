package todo

import (
	"context"

	"github.com/0xirvan/tdl-svelte-go/server/internal/core/domain"
	"github.com/0xirvan/tdl-svelte-go/server/internal/core/port"
)

type GetTodoUsecase struct {
	Repo port.TodoRepository
}

// Execute retrieves a todo item by its ID
func (uc *GetTodoUsecase) Execute(ctx context.Context, id uint) (*domain.Todo, error) {
	return uc.Repo.FindByID(ctx, id)
}
