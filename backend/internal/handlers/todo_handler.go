package handlers

import (
	"strconv"

	"go-backend/internal/models"
	"go-backend/internal/services"

	"github.com/gofiber/fiber/v2"
)

func GetTodos(c *fiber.Ctx) error {
	todos := services.GetAllTodos()

	return c.Status(fiber.StatusOK).JSON(todos)
}

func AddTodo(c *fiber.Ctx) error {
	todo := new(models.Todo)

	if err := c.BodyParser(todo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	if todo.Task == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "[task] can't be empty"})
	}

	services.AddTodo(todo)

	return c.SendStatus(fiber.StatusCreated)
}

func DeleteTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	err = services.DeleteTodoByID(id)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func UpdateTodo(c *fiber.Ctx) error {
	idParam := c.Params("id")

	id, err := strconv.Atoi(idParam)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID format"})
	}

	updatedTodo := new(models.Todo)

	if err := c.BodyParser(updatedTodo); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}

	err = services.UpdateTodoByID(id, updatedTodo)

	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
