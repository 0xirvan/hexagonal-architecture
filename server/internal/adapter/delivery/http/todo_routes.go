package http

import "github.com/labstack/echo/v4"

func RegisterTodoRoutes(e *echo.Echo, h *TodoHandler) {
	g := e.Group("/todos")
	g.GET("/:id", h.GetHandler)
	g.GET("", h.ListHandler)
	g.GET("/paginated", h.ListPaginatedHandler)
	g.POST("", h.CreateHandler)
	g.PATCH("/:id", h.UpdateHandler)
	g.DELETE("/:id", h.DeleteHandler)
	g.PATCH("/:id/done", h.MarkDoneHandler)
	g.PATCH("/:id/undone", h.MarkUndoneHandler)
}
