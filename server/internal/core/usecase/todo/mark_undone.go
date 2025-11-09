package todo

import (
	"context"

	"github.com/0xirvan/tta-svelte-go/server/internal/core/port"
)

type MarkTodoUndoneUsecase struct {
	Repo port.TodoRepository
}

// Execute marks a todo item as not done by its ID, also clearing the CompletedAt timestamp
func (uc *MarkTodoUndoneUsecase) Execute(ctx context.Context, id uint) error {
	todo, err := uc.Repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	todo.IsDone = false
	todo.CompletedAt = nil

	_, err = uc.Repo.Update(ctx, todo)
	return err
}
