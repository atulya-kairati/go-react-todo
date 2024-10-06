package services

import (
	"go-backend/internal/models"
	"go-backend/internal/repositories"
)

func GetAllTodos() []models.Todo {
	return repositories.GetAllTodos()
}

func AddTodo(todo *models.Todo) {
	repositories.AddTodo(todo)
}

func DeleteTodoByID(id int) error {
	return repositories.DeleteTodoByID(id)
}

func UpdateTodoByID(id int, updatedTodo *models.Todo) error {
	return repositories.UpdateTodoByID(id, updatedTodo)
}
