package inmemory

import (
	"context"
	"sync"

	"github.com/0xirvan/tta-svelte-go/server/internal/core/domain"
	"github.com/0xirvan/tta-svelte-go/server/internal/core/port"
)

// TodoRepository implements port.TodoRepository in memory
type TodoRepository struct {
	mu     sync.RWMutex          // mutex for concurrent access
	store  map[uint]*domain.Todo // holds todos in memory using a map
	lastID uint                  // keeps track of the last used ID
}

// NewTodoRepository creates a new instance of TodoRepository
func NewTodoRepository() port.TodoRepository {
	return &TodoRepository{
		store: make(map[uint]*domain.Todo),
	}
}

// Create stores a new todo in memory
func (r *TodoRepository) Create(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	// Increment lastID and assign to the new todo
	r.lastID++
	todo.ID = r.lastID

	r.store[todo.ID] = todo
	return todo, nil
}

// FindByID retrieves a todo by its ID
func (r *TodoRepository) FindByID(ctx context.Context, id uint) (*domain.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	todo, exists := r.store[id]
	if !exists {
		return nil, domain.ErrTodoNotFound
	}

	return todo, nil
}

// FindPaginated retrieves todos in a paginated manner
func (r *TodoRepository) FindPaginated(ctx context.Context, page, size int) ([]*domain.Todo, int, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var todos []*domain.Todo
	for _, todo := range r.store {
		todos = append(todos, todo)
	}

	total := len(todos)

	start := (page - 1) * size
	if start > total {
		return []*domain.Todo{}, total, nil
	}

	end := min(start+size, total)

	return todos[start:end], total, nil
}

// FindAll retrieves all todos
func (r *TodoRepository) FindAll(ctx context.Context) ([]*domain.Todo, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var todos []*domain.Todo
	for _, todo := range r.store {
		todos = append(todos, todo)
	}

	return todos, nil
}

// Delete removes a todo by its ID
func (r *TodoRepository) Delete(ctx context.Context, id uint) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.store[id]; !exists {
		return domain.ErrTodoNotFound
	}

	delete(r.store, id)
	return nil
}

// Update modifies an existing todo
func (r *TodoRepository) Update(ctx context.Context, todo *domain.Todo) (*domain.Todo, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exists := r.store[todo.ID]
	if !exists {
		return nil, domain.ErrTodoNotFound
	}

	r.store[todo.ID] = todo
	return todo, nil
}
