package todo

import (
	"context"
	"time"

	"github.com/0xirvan/tdl-svelte-go/server/internal/core/port"
	"github.com/0xirvan/tdl-svelte-go/server/internal/shared/ptr"
)

type MarkTodoDoneUsecase struct {
	Repo port.TodoRepository
}

// Execute marks a todo item as done by its ID, is also updating the CompletedAt timestamp
func (uc *MarkTodoDoneUsecase) Execute(ctx context.Context, id uint) error {
	todo, err := uc.Repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	now := time.Now()
	todo.IsDone = true
	todo.CompletedAt = ptr.Ptr(now)

	_, err = uc.Repo.Update(ctx, todo)
	return err
}
