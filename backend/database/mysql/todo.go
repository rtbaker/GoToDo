package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	gotodo "github.com/rtbaker/GoToDo/Model"
)

type ToDoService struct {
	db *sql.DB
}

func NewToDoService(db *sql.DB) *ToDoService {
	s := ToDoService{
		db: db,
	}

	return &s
}

func (s *ToDoService) FindToDoByID(ctx context.Context, id int) (*gotodo.ToDo, error) {
	var t gotodo.ToDo

	row := s.db.QueryRowContext(ctx, `SELECT * FROM todo WHERE id = ?`, id)
	if err := row.Scan(&t.ID, &t.UserId, &t.Title, &t.Description, &t.Priority,
		&t.Completed, &t.CreatedAt, &t.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, fmt.Errorf("mysql FindToDoByID %d: %s", id, err)
	}

	return &t, nil
}

// Find a users ToDo's
func (s *ToDoService) FindByUser(ctx context.Context, userId int) ([]*gotodo.ToDo, error) {
	// An albums slice to hold data from returned rows.
	var todos []*gotodo.ToDo

	rows, err := s.db.QueryContext(ctx, `SELECT * FROM todo WHERE userId = ?`, userId)
	if err != nil {
		return nil, fmt.Errorf("mysql FindByUser %d: %s", userId, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var t gotodo.ToDo

		if err := rows.Scan(&t.ID, &t.UserId, &t.Title, &t.Description, &t.Priority,
			&t.Completed, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, fmt.Errorf("mysql FindByUser %d: %s", userId, err)
		}

		todos = append(todos, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("mysql FindByUser %d: %s", userId, err)
	}

	return todos, nil
}

// Create a new ToDo
func (s *ToDoService) CreateToDo(ctx context.Context, todo *gotodo.ToDo) error {
	t := time.Now()
	todo.CreatedAt = t
	todo.UpdatedAt = t

	result, err := s.db.ExecContext(
		ctx,
		`INSERT INTO album (userId, title, description, priority, completed, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?, ?)`,
		todo.UserId,
		todo.Title,
		todo.Description,
		todo.Priority,
		todo.Completed,
		todo.CreatedAt,
		todo.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("mysql CreateToDo: %s", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return fmt.Errorf("mysql CreateToDo: %s", err)
	}

	todo.ID = id

	return nil
}

// Update an existing ToDo
func (s *ToDoService) UpdateToDo(ctx context.Context, id int, upd gotodo.ToDoUpdate) (*gotodo.ToDo, error) {
	t := time.Now()

	result, err := s.db.ExecContext(
		ctx,
		`
		UPDATE todo
		SET title = ?,
		    description = ?,
			priority = ?,
			completed = ?,
			updated_at = ?
		WHERE id = ?
		`,
		upd.Title,
		upd.Description,
		upd.Priority,
		upd.Completed,
		t,
	)

	if err != nil {
		return nil, fmt.Errorf("mysql CreateToDo: %s", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("mysql CreateToDo: %s", err)
	}

	if rows != 1 {
		return nil, fmt.Errorf("mysql CreateToDo: %s", err)
	}

	updatedTodo, err := s.FindToDoByID(ctx, id)

	if err != nil {
		return nil, fmt.Errorf("mysql CreateToDo: %s", err)
	}

	return updatedTodo, nil
}

// Permanently removes a ToDo
func (s *ToDoService) DeleteToDo(ctx context.Context, id int) error {
	// We ignore the result, don't care if it existed or not
	_, err := s.db.ExecContext(
		ctx,
		`DELETE todo WHERE id = ?`,
		id,
	)

	if err != nil {
		return fmt.Errorf("mysql DeleteToDo: %s", err)
	}

	return nil
}
