package http

import (
	"net/http"

	dto "github.com/0xirvan/hexagonal-architecture/server/internal/adapter/delivery/http/dto/todo"
	"github.com/0xirvan/hexagonal-architecture/server/internal/adapter/delivery/http/helper"
	"github.com/0xirvan/hexagonal-architecture/server/internal/core/domain"
	"github.com/0xirvan/hexagonal-architecture/server/internal/core/port"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	todoService port.TodoService
}

func NewTodoHandler(todoService port.TodoService) *TodoHandler {
	return &TodoHandler{todoService: todoService}
}

func (h *TodoHandler) CreateHandler(c echo.Context) error {
	var req dto.CreateTodoRequest

	if err := c.Bind(&req); err != nil {
		return domain.ErrInvalidRequestBody
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	t, err := h.todoService.CreateTodo(c.Request().Context(), req.Title, req.Description)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, dto.ToTodoResponse(t))
}

func (h *TodoHandler) GetHandler(c echo.Context) error {
	id, err := helper.StrToUint(c.Param("id"))
	if err != nil || id <= 0 {
		return domain.ErrInvalidID
	}

	t, err := h.todoService.GetTodo(c.Request().Context(), uint(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.ToTodoResponse(t))
}

func (h *TodoHandler) ListHandler(c echo.Context) error {
	todos, err := h.todoService.ListTodos(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.ToTodoResponseList(todos))
}

func (h *TodoHandler) ListPaginatedHandler(c echo.Context) error {
	page, err := helper.StrToInt(c.QueryParam("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	size, err := helper.StrToInt(c.QueryParam("size"))
	if err != nil || size <= 0 {
		size = 10
	}

	todos, totalItems, err := h.todoService.ListTodosPaginated(c.Request().Context(), page, size)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.ToTodoPaginatedResponse(todos, totalItems, page, size))
}

func (h *TodoHandler) UpdateHandler(c echo.Context) error {
	id, err := helper.StrToUint(c.Param("id"))
	if err != nil || id <= 0 {
		return domain.ErrInvalidID
	}

	var req dto.UpdateTodoRequest
	if err := c.Bind(&req); err != nil {
		return domain.ErrInvalidRequestBody
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	t, err := h.todoService.UpdateTodo(c.Request().Context(), id, req.Title, req.Description)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.ToTodoResponse(t))
}

func (h *TodoHandler) DeleteHandler(c echo.Context) error {
	id, err := helper.StrToUint(c.Param("id"))
	if err != nil || id <= 0 {
		return domain.ErrInvalidID
	}

	if err := h.todoService.DeleteTodo(c.Request().Context(), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *TodoHandler) MarkDoneHandler(c echo.Context) error {
	id, err := helper.StrToUint(c.Param("id"))
	if err != nil || id <= 0 {
		return domain.ErrInvalidID
	}

	t, err := h.todoService.MarkTodoDone(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.ToTodoResponse(t))
}

func (h *TodoHandler) MarkUndoneHandler(c echo.Context) error {
	id, err := helper.StrToUint(c.Param("id"))
	if err != nil || id <= 0 {
		return domain.ErrInvalidID
	}

	t, err := h.todoService.MarkTodoUndone(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.ToTodoResponse(t))
}
