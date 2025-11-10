package port

import (
	"context"

	"github.com/0xirvan/tdl-svelte-go/server/internal/core/domain"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *domain.Todo) (*domain.Todo, error)            // returns created todo
	FindByID(ctx context.Context, id uint) (*domain.Todo, error)                    // returns todo or nil if not found
	FindPaginated(ctx context.Context, page, size int) ([]*domain.Todo, int, error) // returns slice of todos and total count
	FindAll(ctx context.Context) ([]*domain.Todo, error)                            // returns all todos
	Delete(ctx context.Context, id uint) error                                      // deletes a todo, returns error if not found
	Update(ctx context.Context, todo *domain.Todo) (*domain.Todo, error)            // updates a todo, returns updated todo or error if not found
}
