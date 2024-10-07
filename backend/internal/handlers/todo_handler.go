package handlers

import (
	"strconv"

	"go-backend/internal/models"
	"go-backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

type TodoHandler struct {
	Service services.TodoService
}

func NewTodoHandler(service *services.TodoService) *TodoHandler {
	return &TodoHandler{Service: *service}
}

func (th *TodoHandler) GetTodos(c *fiber.Ctx) error {
	todos := th.Service.GetAllTodos()

	return c.Status(fiber.StatusOK).JSON(todos)
}

func (th *TodoHandler) AddTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if todo.Task == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[task] can't be empty"})
	}

	th.Service.AddTodo(todo)

	return c.SendStatus(fiber.StatusCreated)
}

func (th *TodoHandler) DeleteTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	err = th.Service.DeleteTodoByID(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (th *TodoHandler) UpdateTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	updatedTodo := new(models.Todo)

	if err := c.BodyParser(updatedTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	err = th.Service.UpdateTodoByID(id, updatedTodo)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
