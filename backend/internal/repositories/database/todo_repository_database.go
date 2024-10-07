package database

import (
	"log"
	"database/sql"
	"fmt"

	"go-backend/internal/repositories"
	"go-backend/internal/models"
	_ "modernc.org/sqlite"
)

type DatabaseTodoRepository struct {
	db *sql.DB
}

func NewDatabaseTodoRepository() (repositories.TodoRepository, error) {
	db, err := sql.Open("sqlite", "./todos.db")

	if err != nil {
		return nil, err
	}

	createTable(db)

	return &DatabaseTodoRepository{db: db}, nil
}

func (dtr *DatabaseTodoRepository) GetAllTodos() []models.Todo {
	rows, err := dtr.db.Query("select id, task, done from todos");

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	todos := []models.Todo{}

	for rows.Next() {
		todo := new(models.Todo)

		err = rows.Scan(&todo.ID, &todo.Task, &todo.Done)

		if err != nil {
			log.Fatal(err)
		}

		todos = append(todos, *todo)
	}

	return todos
}

func (dtr *DatabaseTodoRepository) AddTodo(todo *models.Todo) {
	insertQuery := `insert into todos (task, done) values (?, ?)`

	_, err := dtr.db.Exec(insertQuery, todo.Task, todo.Done)

	if err != nil {
		log.Fatal(err)
	}

}

func (dtr *DatabaseTodoRepository) DeleteTodoByID(id int) error {
	deleteQuery := `delete from todos where id = ?`

	_, err := dtr.db.Exec(deleteQuery, id)

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func (dtr *DatabaseTodoRepository) UpdateTodoByID(id int, updatedTodo *models.Todo) error {

	if updatedTodo.Task == "" {
		updateQuery := `update todos set done = ? where id = ?`

		_, err := dtr.db.Exec(updateQuery, updatedTodo.Done, id)

		if err != nil {
			log.Fatal(err)
		}
	} else {
		updateQuery := `update todos set task = ?, done = ? where id = ?`

		_, err := dtr.db.Exec(updateQuery, updatedTodo.Task, updatedTodo.Done, id)

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

// creates table if it doesn't exist
func createTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS todos (
		"id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"task" TEXT,
		"done" BOOLEAN
	);`

	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Todos table created successfully or already exists.")
}
