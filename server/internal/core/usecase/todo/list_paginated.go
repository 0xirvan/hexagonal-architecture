package todo

import (
	"context"

	"github.com/0xirvan/tta-svelte-go/server/internal/core/domain"
	"github.com/0xirvan/tta-svelte-go/server/internal/core/port"
)

type ListPaginatedTodosUsecase struct {
	Repo port.TodoRepository
}

// Execute retrieves a paginated list of todo items
func (uc *ListPaginatedTodosUsecase) Execute(ctx context.Context, page, size int) ([]*domain.Todo, int, error) {
	return uc.Repo.FindPaginated(ctx, page, size)
}
