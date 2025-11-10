package todo

import (
	"context"

	"github.com/0xirvan/hexagonal-architecture/server/internal/core/domain"
	"github.com/0xirvan/hexagonal-architecture/server/internal/core/port"
)

// Service implements port.TodoService interface
type Service struct {
	repo port.TodoRepository
}

// NewService creates a new todo service instance
func NewService(repo port.TodoRepository) port.TodoService {
	return &Service{
		repo: repo,
	}
}

// CreateTodo creates a new todo with the given title and description
func (s *Service) CreateTodo(ctx context.Context, title, description string) (*domain.Todo, error) {
	uc := &CreateTodoUsecase{Repo: s.repo}
	return uc.Execute(ctx, title, description)
}

// GetTodo retrieves a todo by its ID
func (s *Service) GetTodo(ctx context.Context, id uint) (*domain.Todo, error) {
	uc := &GetTodoUsecase{Repo: s.repo}
	return uc.Execute(ctx, id)
}

// ListTodos retrieves all todos
func (s *Service) ListTodos(ctx context.Context) ([]*domain.Todo, error) {
	uc := &ListTodosUsecase{Repo: s.repo}
	return uc.Execute(ctx)
}

// ListTodosPaginated retrieves todos with pagination
func (s *Service) ListTodosPaginated(ctx context.Context, page, size int) ([]*domain.Todo, int, error) {
	uc := &ListPaginatedTodosUsecase{Repo: s.repo}
	return uc.Execute(ctx, page, size)
}

// UpdateTodo updates a todo with the given ID
func (s *Service) UpdateTodo(ctx context.Context, id uint, title, description *string) (*domain.Todo, error) {
	uc := &UpdateTodoUsecase{Repo: s.repo}
	return uc.Execute(ctx, id, title, description)
}

// DeleteTodo deletes a todo by its ID
func (s *Service) DeleteTodo(ctx context.Context, id uint) error {
	uc := &DeleteTodoUsecase{Repo: s.repo}
	return uc.Execute(ctx, id)
}

// MarkTodoDone marks a todo as done
func (s *Service) MarkTodoDone(ctx context.Context, id uint) (*domain.Todo, error) {
	uc := &MarkTodoDoneUsecase{Repo: s.repo}
	return uc.Execute(ctx, id)
}

// MarkTodoUndone marks a todo as undone
func (s *Service) MarkTodoUndone(ctx context.Context, id uint) (*domain.Todo, error) {
	uc := &MarkTodoUndoneUsecase{Repo: s.repo}
	return uc.Execute(ctx, id)
}
