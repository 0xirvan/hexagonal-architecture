package http

import (
	"net/http"

	dto "github.com/0xirvan/tdl-svelte-go/server/internal/adapter/delivery/http/dto/todo"
	"github.com/0xirvan/tdl-svelte-go/server/internal/adapter/delivery/http/helper"
	"github.com/0xirvan/tdl-svelte-go/server/internal/core/domain"
	"github.com/0xirvan/tdl-svelte-go/server/internal/core/usecase/todo"

	"github.com/labstack/echo/v4"
)

type TodoHandler struct {
	uc *todo.Service
}

func NewTodoHandler(uc *todo.Service) *TodoHandler {
	return &TodoHandler{uc: uc}
}

func (h *TodoHandler) CreateHandler(c echo.Context) error {
	var req dto.CreateTodoRequest

	if err := c.Bind(&req); err != nil {
		return domain.ErrInvalidRequestBody
	}

	if err := c.Validate(&req); err != nil {
		return err
	}

	t, err := h.uc.Create.Execute(c.Request().Context(), req.Title, req.Description)
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

	t, err := h.uc.Get.Execute(c.Request().Context(), uint(id))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, dto.ToTodoResponse(t))
}

func (h *TodoHandler) ListHandler(c echo.Context) error {
	todos, err := h.uc.List.Execute(c.Request().Context())
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

	todos, totalItems, err := h.uc.ListPaginated.Execute(c.Request().Context(), page, size)
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

	t, err := h.uc.Update.Execute(c.Request().Context(), id, req.Title, req.Description)
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

	if err := h.uc.Delete.Execute(c.Request().Context(), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *TodoHandler) MarkDoneHandler(c echo.Context) error {
	id, err := helper.StrToUint(c.Param("id"))
	if err != nil || id <= 0 {
		return domain.ErrInvalidID
	}

	err = h.uc.Done.Execute(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func (h *TodoHandler) MarkUndoneHandler(c echo.Context) error {
	id, err := helper.StrToUint(c.Param("id"))
	if err != nil || id <= 0 {
		return domain.ErrInvalidID
	}

	err = h.uc.Undone.Execute(c.Request().Context(), id)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}
