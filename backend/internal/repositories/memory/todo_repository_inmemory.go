package memory

import (
	"go-backend/internal/models"
	"go-backend/internal/utils"
	"go-backend/internal/repositories"
)

type InMemoryRepository struct {
	todos []models.Todo
}

func NewInMemoryRepository() repositories.TodoRepository {
	return &InMemoryRepository{todos: []models.Todo{}}
}

func (repo *InMemoryRepository) GetAllTodos() []models.Todo {
	return repo.todos
}

func (repo *InMemoryRepository) AddTodo(todo *models.Todo) {
	(*todo).ID = len(repo.todos)

	repo.todos = append(repo.todos, *todo)
}

func (repo *InMemoryRepository) DeleteTodoByID(id int) error {
	newTodos, err := utils.RemoveTodoById(repo.todos, id)

	if err != nil {
		return err
	}

	repo.todos = newTodos

	return nil
}

func (repo *InMemoryRepository) UpdateTodoByID(id int, updatedTodo *models.Todo) error {
	todo, err := utils.FindTodoById(repo.todos, id)

	if err != nil {
		return err
	}

	if updatedTodo.Task != "" {
		todo.Task = updatedTodo.Task
	}

	todo.Done = updatedTodo.Done
	return nil
}
