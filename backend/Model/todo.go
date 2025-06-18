package gotodo

import (
	"context"
	"time"
)

type ToDo struct {
	ID          int64     `json:"id"`
	UserId      int64     `json:"userId"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Priority    int       `json:"priority"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ToDoUpdate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Priority    int    `json:"priority"`
	Completed   bool   `json:"completed"`
}

// Things we can do to a ToDo, to be implemented by database.
type ToDoService interface {
	// Find a ToDo by id.
	FindToDoByID(ctx context.Context, id int) (*ToDo, error)

	// Find a users ToDo's
	FindByUser(ctx context.Context, userId int) ([]*ToDo, error)

	// Create a new ToDo
	CreateToDo(ctx context.Context, todo *ToDo) error

	// Update an existing ToDo
	UpdateToDo(ctx context.Context, id int, upd ToDoUpdate) (*ToDo, error)

	// Permanently removes a ToDo
	DeleteToDo(ctx context.Context, id int) error
}
