package todo

// CreateTodoRequest represents the payload for creating a new todo item
type CreateTodoRequest struct {
	Title       string `json:"title" validate:"required,min=3"`
	Description string `json:"description,omitempty" validate:"omitempty,max=255"`
}

// UpdateTodoRequest represents the payload for updating an existing todo item
type UpdateTodoRequest struct {
	Title       *string `json:"title" validate:"omitempty,min=3"`
	Description *string `json:"description,omitempty" validate:"omitempty,max=255"`
}
