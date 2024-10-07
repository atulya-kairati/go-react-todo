package repositories

import "go-backend/internal/models"

type TodoRepository interface {
	GetAllTodos() []models.Todo
	AddTodo(todo *models.Todo)
	DeleteTodoByID(id int) error 
	UpdateTodoByID(id int, updatedTodo *models.Todo) error 
}
