package inmemory

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"net/mail"
	"strconv"
	"strings"
	"sync"
	"time"

	gotodo "github.com/rtbaker/GoToDo/Model"
)

type UserService struct {
	nextUserId  int
	userById    map[int]*gotodo.User
	userByEmail map[string]*gotodo.User
	mu          sync.RWMutex
}

func NewUserService() *UserService {
	s := UserService{
		nextUserId:  0,
		userById:    make(map[int]*gotodo.User),
		userByEmail: make(map[string]*gotodo.User),
	}

	return &s
}

/**
 * We expect a record of the format:
 *
 * user,<id>,<email address>,<encrypted password>
 *
 */
func (s *UserService) PreloadDataFromFile(filename string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	data, err := readCSVFile(filename)

	if err != nil {
		return fmt.Errorf("inmemory preload data file read error: %s", err)
	}

	reader := csv.NewReader(bytes.NewReader(data))
	reader.FieldsPerRecord = -1

	for {
		record, err := reader.Read()

		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("inmemory preload data file reading CSV data: %s", err)
		}

		recordType := strings.TrimSpace(record[0])

		if recordType == "user" {
			userId, err := strconv.Atoi(strings.TrimSpace(record[1]))
			if err != nil {
				// ignore
				continue
			}
			email := strings.TrimSpace(record[2])
			password := strings.TrimSpace(record[3])

			_, err = mail.ParseAddress(email) // check it's valid email, ignore the line if not
			if err != nil {
				continue
			}

			now := time.Now()

			user := gotodo.User{
				ID:        int64(userId),
				Email:     email,
				Password:  password,
				CreatedAt: now,
				UpdatedAt: now,
			}

			s.userById[int(user.ID)] = &user
			s.userByEmail[user.Email] = &user

			if s.nextUserId <= int(user.ID) {
				s.nextUserId = int(user.ID) + 1
			}
		}
	}

	return nil
}

func (s *UserService) FindUserByID(ctx context.Context, id int) (*gotodo.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.userById[id], nil
}

func (s *UserService) FindUserByEmail(ctx context.Context, email string) (*gotodo.User, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.userByEmail[email], nil
}

// Creates a new user.
func (s *UserService) CreateUser(ctx context.Context, user *gotodo.User) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	t := time.Now()
	user.CreatedAt = t
	user.UpdatedAt = t

	if _, present := s.userByEmail[user.Email]; present {
		return fmt.Errorf("inmemory user with this email already exists")
	}

	s.nextUserId++
	user.ID = int64(s.nextUserId)

	s.userByEmail[user.Email] = user
	s.userById[int(user.ID)] = user

	return nil
}

// Updates a user object.
func (s *UserService) UpdateUserPassword(ctx context.Context, id int, upd gotodo.PasswordUpdate) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.userById[id]
	if !ok {
		return fmt.Errorf("inmemory: user not present in update")
	}

	user.UpdatedAt = time.Now()
	user.Password = upd.NewPassword

	return nil
}

// Permanently deletes a user (and all their ToDo's).
func (s *UserService) DeleteUser(ctx context.Context, id int) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	user, ok := s.userById[id]

	if ok {
		delete(s.userByEmail, user.Email)
		delete(s.userById, id)
	}

	return nil
}
