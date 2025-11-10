package persistence

import (
	"github.com/0xirvan/tdl-svelte-go/server/internal/adapter/persistence/inmemory"
	"github.com/0xirvan/tdl-svelte-go/server/internal/core/port"
)

// for wire stuff
func NewTodoRepository() port.TodoRepository {
	return inmemory.NewTodoRepository()
}
