package repository

import (
	"todo-app/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	CreateTodo(todo models.Todo) (models.Todo, error)
	GetTodo(TodoId int) (models.Todo, error)
	FindTodoByGroupId(GroupID int) ([]models.Todo, error)
	FindTodo() ([]models.Todo, error)
	UpdateTodo(todo models.Todo) (models.Todo, error)
	DeleteTodo(todo models.Todo) (models.Todo, error)
}

func RepositoryTodo(db *gorm.DB) *repository {
	return &repository{db}
}
func (r *repository) UpdateTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Save(&todo).Error
	return todo, err
}

func (r *repository) DeleteTodo(todo models.Todo) (models.Todo, error) {
	err := r.db.Delete(&todo).Error // Using Delete method

	return todo, err
}

func (r *repository) CreateTodo(Todo models.Todo) (models.Todo, error) {
	err := r.db.Create(&Todo).Error
	return Todo, err
}

func (r *repository) GetTodo(TodoId int) (models.Todo, error) {
	var Todo models.Todo
	err := r.db.Preload("ActivityGroup").First(&Todo, TodoId).Error
	return Todo, err

}

func (r *repository) FindTodo() ([]models.Todo, error) {
	var Todo []models.Todo
	err := r.db.Preload("ActivityGroup").Raw("SELECT * FROM todos").Scan(&Todo).Error

	return Todo, err
}

func (r *repository) FindTodoByGroupId(GroupID int) ([]models.Todo, error) {
	var todo []models.Todo
	err := r.db.Preload("ActivityGroup").Where("activity_group_id = ?", GroupID).Find(&todo).Error

	return todo, err
}
