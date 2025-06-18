package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	gotodo "github.com/rtbaker/GoToDo/Model"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	s := UserService{
		db: db,
	}

	return &s
}

func (s *UserService) FindUserByID(ctx context.Context, id int) (*gotodo.User, error) {
	var u gotodo.User

	row := s.db.QueryRowContext(ctx, `SELECT * FROM users WHERE id = ?`, id)
	if err := row.Scan(&u.ID, &u.Email, &u.Password, &u.CreatedAt, &u.UpdatedAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}

		return nil, fmt.Errorf("mysql FindUserByID %d: %s", id, err)
	}

	return &u, nil
}

// Creates a new user.
func (s *UserService) CreateUser(ctx context.Context, user *gotodo.User) error {
	t := time.Now()
	user.CreatedAt = t
	user.UpdatedAt = t

	result, err := s.db.ExecContext(
		ctx,
		`INSERT INTO users (email, password, created_at, updated_at) VALUES (?, ?, ?, ?)`,
		user.Email,
		user.Password,
		user.CreatedAt,
		user.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("mysql CreateUser: %s", err)
	}

	id, err := result.LastInsertId()

	if err != nil {
		return fmt.Errorf("mysql CreateUser: %s", err)
	}

	user.ID = id

	return nil
}

// Updates a user object.
func (s *UserService) UpdateUserPassword(ctx context.Context, id int, upd gotodo.PasswordUpdate) error {
	t := time.Now()

	result, err := s.db.ExecContext(
		ctx,
		`
		UPDATE users
		SET password = ?,
		WHERE id = ?
		`,
		upd.NewPassword,
		t,
	)

	if err != nil {
		return fmt.Errorf("mysql UpdateUserPassword: %s", err)
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("mysql UpdateUserPassword: %s", err)
	}

	if rows != 1 {
		return fmt.Errorf("mysql UpdateUserPassword: %s", err)
	}

	return nil
}

// Permanently deletes a user (and all their ToDo's).
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	// We ignore the result, don't care if it existed or not
	_, err := s.db.ExecContext(
		ctx,
		`DELETE FROM users WHERE id = ?`,
		id,
	)

	if err != nil {
		return fmt.Errorf("mysql DeleteUser: %s", err)
	}

	return nil
}
