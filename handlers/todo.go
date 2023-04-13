package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	dto "todo-app/dto/result"
	tododto "todo-app/dto/todo"
	"todo-app/models"
	"todo-app/repository"

	"github.com/labstack/echo/v4"
)

type handlerTodo struct {
	TodoRepository repository.TodoRepository
}

func HandlerTodo(TodoRepository repository.TodoRepository) *handlerTodo {
	return &handlerTodo{TodoRepository}
}

func (h *handlerTodo) UpdateTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	request := new(tododto.UpdateTodoRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	todo, err := h.TodoRepository.GetTodo(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResultCusErr{Status: ResultStatus(http.StatusBadRequest), Message: ResultMessage(err, id)})
	}

	if err != nil {
		fmt.Println("Error:", err)
	}

	if request.Title != "" {
		todo.Title = request.Title
	}
	if request.Priority != "" {
		todo.Priority = request.Priority
	}
	todo.Status = request.Status
	if request.Status != "" {
		todo.Status = request.Status
	}
	todo.UpdateAt = time.Now()

	todos := models.Todo{
		ID:              todo.ID,
		ActivityGroupId: todo.ActivityGroupId,
		Title:           todo.Title,
		Priority:        todo.Priority,
		Status:          todo.Status,
		CreatedAt:       todo.CreatedAt,
		UpdateAt:        todo.UpdateAt,
	}

	data, err := h.TodoRepository.UpdateTodo(todos)
	if err != nil {

		return c.JSON(http.StatusBadRequest, dto.ResultCusErr{Status: ResultStatus(http.StatusBadRequest), Message: ResultMessage(err, id)})
	}

	return c.JSON(http.StatusOK, dto.ResultCusSucces{Status: ResultStatus(http.StatusOK), Message: ResultStatus(http.StatusOK), Data: convertTodoRes(data)})

}

func (h *handlerTodo) DeleteTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.TodoRepository.GetTodo(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResultCusErr{Status: ResultStatus(http.StatusBadRequest), Message: ResultMessage(err, id)})
	}

	deleteTodo, err := h.TodoRepository.DeleteTodo(todo)
	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResultCusErr{Status: ResultStatus(http.StatusBadRequest), Message: ResultMessage(err, id)})
	}
	return c.JSON(http.StatusOK, dto.ResultCusSucces{Status: ResultStatus(http.StatusOK), Message: ResultStatus(http.StatusOK), Data: convertTodoRes(deleteTodo)})
}

func (h *handlerTodo) GetTodo(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.TodoRepository.GetTodo(id)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dto.ResultCusErr{Status: ResultStatus(http.StatusBadRequest), Message: ResultMessage(err, id)})
	}

	return c.JSON(http.StatusOK, dto.ResultCusSucces{Status: ResultStatus(http.StatusOK), Message: ResultStatus(http.StatusOK), Data: convertTodoRes(todo)})

}

func (h *handlerTodo) FindTodo(c echo.Context) error {
	todo, err := h.TodoRepository.FindTodo()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}

	return c.JSON(http.StatusOK, dto.ResultCusSucces{Status: ResultStatus(http.StatusOK), Message: ResultStatus(http.StatusOK), Data: convertAllTodo(todo)})

}

func (h *handlerTodo) FindTodoByGroupId(c echo.Context) error {
	idActivityGroup, _ := strconv.Atoi(c.Param("id"))
	todo, err := h.TodoRepository.FindTodoByGroupId(idActivityGroup)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.ResultCusSucces{Status: ResultStatus(http.StatusOK), Message: ResultStatus(http.StatusOK), Data: convertAllTodo(todo)})

}

func (h *handlerTodo) CreateTodo(c echo.Context) error {
	request := new(tododto.CreateTodoRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusBadRequest, dto.ErrorResult{Code: http.StatusBadRequest, Message: err.Error()})
	}

	var pr string

	if request.Priority != "" {
		pr = request.Priority
	} else {
		pr = "very-high"
	}

	todo := models.Todo{
		Title:           request.Title,
		ActivityGroupId: request.ActivityGroupId,
		IsActive:        request.IsActive,
		Priority:        pr,
		CreatedAt:       time.Now(),
		UpdateAt:        time.Now(),
	}

	newTodo, err := h.TodoRepository.CreateTodo(todo)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dto.ErrorResult{Code: http.StatusInternalServerError, Message: err.Error()})
	}
	return c.JSON(http.StatusOK, dto.ResultCusSucces{Status: ResultStatus(http.StatusOK), Message: ResultStatus(http.StatusOK), Data: convertTodoRes(newTodo)})

}

func convertAllTodo(u []models.Todo) []tododto.TodoResponse {
	responses := make([]tododto.TodoResponse, 0)
	for _, todo := range u {
		response := tododto.TodoResponse{
			ID:              todo.ID,
			ActivityGroupId: todo.ActivityGroupId,
			Title:           todo.Title,
			IsActive:        todo.IsActive,
			Priority:        todo.Priority,
			CreatedAt:       todo.CreatedAt,
			UpdateAt:        todo.UpdateAt,
		}
		responses = append(responses, response)
	}
	return responses
}

func convertTodoRes(u models.Todo) tododto.TodoResponse {
	return tododto.TodoResponse{
		ID:              u.ID,
		Title:           u.Title,
		ActivityGroupId: u.ActivityGroupId,
		IsActive:        u.IsActive,
		Priority:        u.Priority,
		CreatedAt:       u.CreatedAt,
		UpdateAt:        u.UpdateAt,
	}
}
