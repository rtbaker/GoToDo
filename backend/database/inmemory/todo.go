package inmemory

import (
	"bytes"
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"

	gotodo "github.com/rtbaker/GoToDo/Model"
)

type ToDoService struct {
	nextTodoId    int
	totdosById    map[int]*gotodo.ToDo
	todosByUserId map[int][]*gotodo.ToDo
	mu            sync.RWMutex
}

func NewToDoService() *ToDoService {
	s := ToDoService{
		nextTodoId:    0,
		totdosById:    make(map[int]*gotodo.ToDo),
		todosByUserId: make(map[int][]*gotodo.ToDo),
	}

	return &s
}

/**
 * We expect a record of the format:
 *
 * todo,<userId>,<title>,<description>,<priority>,<completed>
 *
 */
func (s *ToDoService) PreloadDataFromFile(filename string) error {
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

		if recordType == "todo" {
			userId, err := strconv.Atoi(strings.TrimSpace(record[1]))
			if err != nil {
				// ignore
				continue
			}

			title := strings.TrimSpace(record[2])
			description := strings.TrimSpace(record[3])

			priority, err := strconv.Atoi(strings.TrimSpace(record[4]))
			if err != nil {
				// ignore
				continue
			}

			completed := false
			if strings.TrimSpace(record[4]) == "true" {
				completed = true
			}

			now := time.Now()

			todo := gotodo.ToDo{
				ID:          int64(s.nextTodoId),
				UserId:      int64(userId),
				Title:       title,
				Description: description,
				Priority:    priority,
				Completed:   completed,
				UpdatedAt:   now,
				CreatedAt:   now,
			}

			s.totdosById[int(todo.ID)] = &todo

			_, ok := s.todosByUserId[int(todo.UserId)]

			if !ok {
				s.todosByUserId[int(todo.UserId)] = make([]*gotodo.ToDo, 0)
			}

			s.todosByUserId[int(todo.UserId)] = append(s.todosByUserId[int(todo.UserId)], &todo)
			s.nextTodoId++

		}
	}

	return nil
}

func (s *ToDoService) FindToDoByID(ctx context.Context, id int) (*gotodo.ToDo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.totdosById[id], nil
}

// Find a users ToDo's
func (s *ToDoService) FindByUser(ctx context.Context, userId int) ([]*gotodo.ToDo, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.todosByUserId[userId], nil
}

// Create a new ToDo
func (s *ToDoService) CreateToDo(ctx context.Context, todo *gotodo.ToDo) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	t := time.Now()
	todo.CreatedAt = t
	todo.UpdatedAt = t

	s.nextTodoId++

	todo.ID = int64(s.nextTodoId)

	s.todosByUserId[int(todo.UserId)] = append(s.todosByUserId[int(todo.UserId)], todo)
	s.totdosById[int(todo.ID)] = todo

	return nil
}

// Update an existing ToDo
func (s *ToDoService) UpdateToDo(ctx context.Context, id int, upd gotodo.ToDoUpdate) (*gotodo.ToDo, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.totdosById[id]
	if !ok {
		return nil, fmt.Errorf("inmemory: todo not present in update")
	}

	todo.UpdatedAt = time.Now()
	todo.Title = upd.Title
	todo.Description = upd.Description
	todo.Priority = upd.Priority
	todo.Completed = upd.Completed

	return todo, nil
}

// Permanently removes a ToDo
func (s *ToDoService) DeleteToDo(ctx context.Context, id int) error {
	// We ignore the result, don't care if it existed or not
	s.mu.Lock()
	defer s.mu.Unlock()

	todo, ok := s.totdosById[id]

	if ok {
		delete(s.totdosById, id)

		index := slices.Index(s.todosByUserId[int(todo.UserId)], todo)
		s.todosByUserId[int(todo.UserId)] = slices.Delete(s.todosByUserId[int(todo.UserId)], index, index+1)
	}

	return nil
}
