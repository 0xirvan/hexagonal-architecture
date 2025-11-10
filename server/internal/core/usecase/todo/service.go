package todo

import "github.com/0xirvan/tdl-svelte-go/server/internal/core/port"

type Service struct {
	Create        *CreateTodoUsecase
	Get           *GetTodoUsecase
	List          *ListTodosUsecase
	ListPaginated *ListPaginatedTodosUsecase
	Update        *UpdateTodoUsecase
	Delete        *DeleteTodoUsecase
	Done          *MarkTodoDoneUsecase
	Undone        *MarkTodoUndoneUsecase
}

func NewService(repo port.TodoRepository) *Service {
	return &Service{
		Create:        &CreateTodoUsecase{Repo: repo},
		Get:           &GetTodoUsecase{Repo: repo},
		List:          &ListTodosUsecase{Repo: repo},
		ListPaginated: &ListPaginatedTodosUsecase{Repo: repo},
		Update:        &UpdateTodoUsecase{Repo: repo},
		Delete:        &DeleteTodoUsecase{Repo: repo},
		Done:          &MarkTodoDoneUsecase{Repo: repo},
		Undone:        &MarkTodoUndoneUsecase{Repo: repo},
	}
}
