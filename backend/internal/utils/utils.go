package utils

import (
	"errors"
	"fmt"

	"go-backend/internal/models"
)

func RemoveTodoById(todos []models.Todo, id int) ([]models.Todo, error) {

	for i, todo := range todos {
		if todo.ID == id {
			return append(todos[:i], todos[i+1:]...), nil
		}
	}

	return todos, errors.New(fmt.Sprintf("Todo with id:%d doesn't exist", id))
}

func FindTodoById(todos []models.Todo, id int) (*models.Todo, error) {
	for i := range todos {
		if todos[i].ID == id {
			return &todos[i], nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Todo with id:%d doesn't exist", id))
}
