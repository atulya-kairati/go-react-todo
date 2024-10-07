package services

import (
	"go-backend/internal/models"
	"go-backend/internal/repositories"
)

type TodoService struct {
	Repo repositories.TodoRepository
}

func NewTodoService(repository repositories.TodoRepository) *TodoService {
	return &TodoService{Repo: repository}
}

func (service *TodoService) GetAllTodos() []models.Todo {
	return service.Repo.GetAllTodos()
}

func (service *TodoService) AddTodo(todo *models.Todo) {
	service.Repo.AddTodo(todo)
}

func (service *TodoService) DeleteTodoByID(id int) error {
	return service.Repo.DeleteTodoByID(id)
}

func (service *TodoService) UpdateTodoByID(id int, updatedTodo *models.Todo) error {
	return service.Repo.UpdateTodoByID(id, updatedTodo)
}
