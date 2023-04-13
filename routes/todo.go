package routes

import (
	"todo-app/handlers"
	"todo-app/pkg/mysql"
	"todo-app/repository"

	"github.com/labstack/echo/v4"
)

func TodoRoutes(e *echo.Group) {
	todoRepository := repository.RepositoryTodo(mysql.DB)
	h := handlers.HandlerTodo(todoRepository)

	e.POST("/todo-items", h.CreateTodo)
	e.GET("/todo-items", h.FindTodo)
	e.GET("/todo-items/:id", h.GetTodo)
	e.DELETE("/todo-items/:id", h.DeleteTodo)
	e.GET("/todos-items/:id", h.FindTodoByGroupId)
	e.PATCH("/todo-items/:id", h.UpdateTodo)

}
