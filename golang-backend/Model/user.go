package gotodo

import (
	"context"
	"encoding/gob"
	"time"
)

// SCS session uses gob, we need to register our user struct so we can store it in
// the session data
func init() {
	gob.Register(User{})
}

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type PasswordUpdate struct {
	NewPassword string `json:"newPassword"`
	OldPassword string `json:"updatePassword"`
}

// Things we can do to a user, to be implemented by database.
type UserService interface {
	// Get a user by id
	FindUserByID(ctx context.Context, id int) (*User, error)

	// Get a user by email address
	FindUserByEmail(ctx context.Context, email string) (*User, error)

	// Retrieves a list of users by filter. Also returns total count of matching
	// users which may differ from returned results if filter.Limit is specified.
	//FindUsers(ctx context.Context, filter UserFilter) ([]*User, int, error)

	// Creates a new user.
	CreateUser(ctx context.Context, user *User) error

	// Updates a user object.
	UpdateUserPassword(ctx context.Context, id int, upd PasswordUpdate) error

	// Permanently deletes a user (and all their ToDo's).
	DeleteUser(ctx context.Context, id int) error
}
