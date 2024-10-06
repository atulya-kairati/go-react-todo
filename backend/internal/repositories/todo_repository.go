package repositories

import (
	"go-backend/internal/models"
	"go-backend/internal/utils"
)

var todos []models.Todo

func init() {
	todos = []models.Todo{}
}

func GetAllTodos() []models.Todo {
	return todos
}

func AddTodo(todo *models.Todo) {
	(*todo).ID = len(todos)

	todos = append(todos, *todo)
}

func DeleteTodoByID(id int) error {
	newTodos, err := utils.RemoveTodoById(todos, id)

	if err != nil {
		return err
	}

	todos = newTodos

	return nil
}

func UpdateTodoByID(id int, updatedTodo *models.Todo) error {
	todo, err := utils.FindTodoById(todos, id)

	if err != nil {
		return err
	}

	if updatedTodo.Task != "" {
		todo.Task = updatedTodo.Task
	}

	todo.Done = updatedTodo.Done
	return nil
}
