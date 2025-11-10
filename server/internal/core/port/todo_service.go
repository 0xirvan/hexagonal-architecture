package port

import (
	"context"

	"github.com/0xirvan/hexagonal-architecture/server/internal/core/domain"
)

// TodoService defines the business logic operations for todos
type TodoService interface {
	// CreateTodo creates a new todo with the given title and description
	CreateTodo(ctx context.Context, title, description string) (*domain.Todo, error)

	// GetTodo retrieves a todo by its ID
	GetTodo(ctx context.Context, id uint) (*domain.Todo, error)

	// ListTodos retrieves all todos
	ListTodos(ctx context.Context) ([]*domain.Todo, error)

	// ListTodosPaginated retrieves todos with pagination
	ListTodosPaginated(ctx context.Context, page, size int) ([]*domain.Todo, int, error)

	// UpdateTodo updates a todo with the given ID
	UpdateTodo(ctx context.Context, id uint, title, description *string) (*domain.Todo, error)

	// DeleteTodo deletes a todo by its ID
	DeleteTodo(ctx context.Context, id uint) error

	// MarkTodoDone marks a todo as done
	MarkTodoDone(ctx context.Context, id uint) (*domain.Todo, error)

	// MarkTodoUndone marks a todo as undone
	MarkTodoUndone(ctx context.Context, id uint) (*domain.Todo, error)
}
