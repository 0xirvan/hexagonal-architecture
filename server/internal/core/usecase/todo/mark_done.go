package todo

import (
	"context"
	"time"

	"github.com/0xirvan/hexagonal-architecture/server/internal/core/domain"
	"github.com/0xirvan/hexagonal-architecture/server/internal/core/port"
	"github.com/0xirvan/hexagonal-architecture/server/internal/shared/ptr"
)

type MarkTodoDoneUsecase struct {
	Repo port.TodoRepository
}

// Execute marks a todo item as done by its ID, is also updating the CompletedAt timestamp
func (uc *MarkTodoDoneUsecase) Execute(ctx context.Context, id uint) (*domain.Todo, error) {
	todo, err := uc.Repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	now := time.Now()
	todo.IsDone = true
	todo.CompletedAt = ptr.Ptr(now)

	return uc.Repo.Update(ctx, todo)
}
