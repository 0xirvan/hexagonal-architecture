package todo

import (
	"time"

	"github.com/0xirvan/hexagonal-architecture/server/internal/adapter/delivery/http/dto"
	"github.com/0xirvan/hexagonal-architecture/server/internal/core/domain"
	"github.com/0xirvan/hexagonal-architecture/server/internal/shared/ptr"
)

// TodoResponse represents the response payload for a todo item
type TodoResponse struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	IsDone      bool    `json:"is_done"`
	CompletedAt *string `json:"completed_at"`
	CreatedAt   string  `json:"created_at"`
}

// TodoPaginatedResponse represents a paginated response of todo items
type TodoPaginatedResponse = dto.PaginatedResponse[TodoResponse]

// ToTodoResponse maps a domain.Todo to a TodoResponse
func ToTodoResponse(t *domain.Todo) TodoResponse {
	var completed *string
	if t.CompletedAt != nil {
		s := t.CompletedAt.Format(time.RFC3339)
		completed = ptr.Ptr(s)
	}

	return TodoResponse{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		IsDone:      t.IsDone,
		CompletedAt: completed,
		CreatedAt:   t.CreatedAt.Format(time.RFC3339),
	}
}

// ToTodoResponseList maps a list of domain.Todo to a list of TodoResponse
func ToTodoResponseList(list []*domain.Todo) []TodoResponse {
	out := make([]TodoResponse, len(list))
	for i, t := range list {
		out[i] = ToTodoResponse(t)
	}
	return out
}

// ToTodoPaginatedResponse converts list of domain.Todo + metadata to TodoPaginatedResponse
func ToTodoPaginatedResponse(
	todos []*domain.Todo,
	totalItems int,
	page int,
	pageSize int,
) TodoPaginatedResponse {
	totalPages := (totalItems + pageSize - 1) / pageSize

	return TodoPaginatedResponse{
		Data:        ToTodoResponseList(todos),
		TotalItems:  totalItems,
		TotalPages:  totalPages,
		CurrentPage: page,
		PageSize:    pageSize,
		HasNext:     page < totalPages,
		HasPrev:     page > 1,
	}
}
