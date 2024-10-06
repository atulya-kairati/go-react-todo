package main

import (
	"fmt"
	"log"
	"os"

	"go-backend/internal/handlers"
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

	app := fiber.New()
	app.Use(logger.New())

	app.Get("/todos", handlers.GetTodos)

	app.Post("/todo", handlers.AddTodo)

	app.Delete("/todo/:id", handlers.DeleteTodo)

	app.Patch("/todo/:id", handlers.UpdateTodo)

	app.Listen(":" + PORT)

}
