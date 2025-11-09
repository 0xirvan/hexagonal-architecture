package todo

import (
	"context"

	"github.com/0xirvan/tta-svelte-go/server/internal/core/port"
)

type DeleteTodoUsecase struct {
	Repo port.TodoRepository
}

// Execute deletes a todo item by its ID
func (uc *DeleteTodoUsecase) Execute(ctx context.Context, id uint) error {
	return uc.Repo.Delete(ctx, id)
}
