package todo

import (
	"context"
	"time"

	"github.com/0xirvan/tdl-svelte-go/server/internal/core/domain"
	"github.com/0xirvan/tdl-svelte-go/server/internal/core/port"
)

type CreateTodoUsecase struct {
	Repo port.TodoRepository
}

// Execute creates a new todo item with the given title and description
func (uc *CreateTodoUsecase) Execute(ctx context.Context, title, desc string) (*domain.Todo, error) {
	todo := &domain.Todo{
		Title:       title,
		Description: desc,
		CreatedAt:   time.Now(),
	}

	return uc.Repo.Create(ctx, todo)
}
