package todoservice

import (
	"fmt"
	"log"
	"time"

	dtos "github.com/simonteh28/go-todo-app/api/dtos/todo"
	"github.com/simonteh28/go-todo-app/config"
	"github.com/simonteh28/go-todo-app/internal/db"
)

type TodoService interface {
	Post(dtos.Todo) (int64, error)
	Get() ([]*dtos.Todo, error)
	GetById(string) (*dtos.Todo, error)
	Patch(string, dtos.UpdateTodo) (string, error)
	Delete(string) error
}

type TodoHandler struct {
	db *db.DB
}

// Instantiates a new instance of Todo Service
func New(cfg *config.Config) (*TodoHandler, error) {
	th := &TodoHandler{}
	var err error
	
	// Initialize database
	th.db, err = db.Get(cfg.GetDBConnString())
	if err != nil {
		return nil, fmt.Errorf("could not connect to db: %w", err)
	}		
	log.Println("Successfully connected to db: " + cfg.GetDBString())

	return th, nil
}

func (th *TodoHandler) Post(todo dtos.Todo) (int64, error) {

	// log.Println("Inserting new record: " + todo.Title)
	_, err := th.db.Client.Exec("INSERT INTO todo(title, description, completed, date) VALUES ($1, $2, $3, $4)", todo.Title, todo.Description, false, todo.Date)
	if err != nil {
		return 0, fmt.Errorf("unable to execute insert: %v", err)
	}
	
	// Query the last inserted record
	row := th.db.Client.QueryRow("SELECT LASTVAL()")
	err = row.Scan(&todo.ID)
	if err != nil {
		return 0, fmt.Errorf("unable to get the last inserted row: %v", err)
	}
	
    return todo.ID, nil
}

// Gets all record of todo
func (th *TodoHandler) Get() ([]*dtos.Todo, error) {
	var todos []*dtos.Todo

	// Perform select all
	rows, err := th.db.Client.Query("SELECT id, title, description, completed, date FROM todo")
	if err != nil {
        return nil, fmt.Errorf("error fetching all todos: %v", err)
    }
	defer rows.Close()
	
	// Loop rows
	for rows.Next() {
		var todo dtos.Todo
		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.Date); err != nil {
			return nil, fmt.Errorf("error fetching all todos: %v", err)
		}
		todos = append(todos, &todo)
	}
	
	// If contains error throw
	if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("error fetching all todos: %v", err)
    }

	return todos, nil
}

// Get todo by id
func (th *TodoHandler) GetById(id string) (*dtos.Todo, error){
	var todo dtos.Todo
	// Handle connection to db and perform select with condition
	row := th.db.Client.QueryRow("SELECT id, title, description, completed, updated_at, date FROM todo WHERE id = $1", id)
	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Completed, &todo.UpdatedDate, &todo.Date)
	if err != nil {
		return nil, fmt.Errorf("unable to get the todo with id " + id + "error: %v", err)
	}

	return &todo, nil
}

func (th *TodoHandler) Patch(id string, updateTodo dtos.UpdateTodo) (string, error) {
	var todo *dtos.Todo

	todo, err := th.GetById(id)
	if err != nil {
		return "Failed", err
	}
	
	// Map field changes
	todo.MapChanges(updateTodo)

	// Handle connection to db and perform update with condition
	_, err = th.db.Client.Exec("UPDATE todo SET title = $1, description = $2, completed = $3, updated_at = $4, date = $5 WHERE id = $6", todo.Title, todo.Description, todo.Completed, time.Now(), todo.Date, id)
	if err != nil {
		return "Failed", fmt.Errorf("unable to update the todo with id " + id + "error: %v", err)
	}

	return "Updated todo " + todo.Title + " successfuly", nil
}

func (th *TodoHandler) Delete(id string) error {

	// Handle connection to db and perform delete from
	_, err := th.db.Client.Exec("DELETE FROM todo WHERE id = $1", id)
	if err != nil {
		return fmt.Errorf("unable to delete the todo with id " + id + "error: %v", err)
	}

	return nil
}
