package todo

// CreateTodoRequest represents the payload for creating a new todo item
type CreateTodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

// UpdateTodoRequest represents the payload for updating an existing todo item
type UpdateTodoRequest struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

// MarkTodoRequest represents the payload for marking a todo item as done or undone
type MarkTodoRequest struct {
	IsDone bool `json:"is_done"`
}
