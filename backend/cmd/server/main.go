package main

import (
	"fmt"
	"log"
	"os"

	"go-backend/internal/handlers"
	"go-backend/internal/repositories"
	"go-backend/internal/services"
	"go-backend/internal/repositories/memory"
	"go-backend/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
)

var todos []models.Todo

func main() {

	todos = []models.Todo{}

	fmt.Println("Starting server...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	repoType := os.Getenv("REPO")

	app := fiber.New()
	app.Use(logger.New())

	// create handler
	repo := getRepository(repoType)
	service := services.NewTodoService(repo)
	handler := handlers.NewTodoHandler(service)

	app.Get("/todos", handler.GetTodos)

	app.Post("/todo", handler.AddTodo)

	app.Delete("/todo/:id", handler.DeleteTodo)

	app.Patch("/todo/:id", handler.UpdateTodo)

	app.Listen(":" + PORT)

}

func getRepository(repoType string) repositories.TodoRepository {
	if repoType == "MEMORY" {
		return memory.NewInMemoryRepository()
	}
	return nil
}
