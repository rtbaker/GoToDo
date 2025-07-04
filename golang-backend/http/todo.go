package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	gotodo "github.com/rtbaker/GoToDo/Model"
)

func (s *Server) registerTodoRoutes(r *mux.Router) {
	// Listing of all todos for a user.
	r.HandleFunc("/api/1.0/todos", s.handleTodoIndex).Methods("GET")
	r.HandleFunc("/api/1.0/todos", s.handleCreateTodo).Methods("POST")
	r.HandleFunc("/api/1.0/todos/{id:[0-9]+}", s.handleDeleteTodo).Methods("DELETE")
	r.HandleFunc("/api/1.0/todos/{id:[0-9]+}", s.handleUpdateTodo).Methods("PATCH")
}

func (s *Server) handleTodoIndex(w http.ResponseWriter, r *http.Request) {
	user, ok := s.getUser(w, r)
	if !ok {
		return
	}

	todos, err := s.TodoService.FindByUser(r.Context(), int(user.ID))

	if err != nil {
		httpErr := HttpError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("%s", err),
		}

		ReturnError(w, httpErr)
	}

	ReturnJson(w, http.StatusOK, todos)
}

type CreateTodoDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
	Priority    int    `json:"priority" validate:"required"`
}

func (s *Server) handleCreateTodo(w http.ResponseWriter, r *http.Request) {
	user, ok := s.getUser(w, r)
	if !ok {
		return
	}

	var newTodo CreateTodoDTO

	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		ReturnError(w, HttpError{http.StatusBadRequest, "invalid JSON body"})
		return
	}

	if err := s.validator.Struct(newTodo); err != nil {
		ReturnError(w, HttpError{http.StatusBadRequest, fmt.Sprintf("validation error: %s", err)})
		return
	}

	todo := gotodo.ToDo{
		UserId:      user.ID,
		Title:       newTodo.Title,
		Description: newTodo.Description,
		Priority:    newTodo.Priority,
	}

	if err := s.TodoService.CreateToDo(r.Context(), &todo); err != nil {
		ReturnError(w, HttpError{http.StatusInternalServerError, fmt.Sprintf("db error: %s", err)})
		return
	}

	ReturnJson(w, http.StatusOK, todo)
}

func (s *Server) handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	todo, ok := s.getRequestTodoById(w, r)

	if !ok {
		return
	}

	err := s.TodoService.DeleteToDo(r.Context(), int(todo.ID))

	if err != nil {
		ReturnError(w, HttpError{http.StatusInternalServerError, fmt.Sprintf("db error: %s", err)})
		return
	}

	w.WriteHeader(http.StatusOK)
}

type UpdateTodoDTO struct {
	Title       *string `json:"title"`
	Description *string `json:"description"`
	Priority    *int    `json:"priority"`
	Completed   *bool   `json:"completed"`
}

func (s *Server) handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	todo, ok := s.getRequestTodoById(w, r)

	if !ok {
		return
	}

	var updateTodo UpdateTodoDTO

	if err := json.NewDecoder(r.Body).Decode(&updateTodo); err != nil {
		ReturnError(w, HttpError{http.StatusBadRequest, "invalid JSON body"})
		return
	}

	// there is probably a better way of doing this
	var changed bool

	if updateTodo.Title != nil {
		todo.Title = *updateTodo.Title
		changed = true
	}

	if updateTodo.Description != nil {
		todo.Description = *updateTodo.Description
		changed = true
	}

	if updateTodo.Priority != nil {
		todo.Priority = *updateTodo.Priority
		changed = true
	}

	if updateTodo.Completed != nil {
		todo.Completed = *updateTodo.Completed
		changed = true
	}

	if changed {
		todoUp := gotodo.ToDoUpdate{
			Title:       todo.Title,
			Description: todo.Description,
			Priority:    todo.Priority,
			Completed:   todo.Completed,
		}

		var err error
		todo, err = s.TodoService.UpdateToDo(r.Context(), int(todo.ID), todoUp)

		if err != nil {
			ReturnError(w, HttpError{http.StatusInternalServerError, fmt.Sprintf("db error: %s", err)})
			return
		}
	}

	ReturnJson(w, http.StatusOK, todo)
}

// Get the todo specified in the "id" url parameter and check it belongs to the user
// send errors if not ok
func (s *Server) getRequestTodoById(w http.ResponseWriter, r *http.Request) (*gotodo.ToDo, bool) {
	user, ok := s.getUser(w, r)
	if !ok {
		return nil, false
	}

	vars := mux.Vars(r)
	todoId, err := strconv.Atoi(vars["id"])

	if err != nil {
		ReturnError(w, HttpError{http.StatusInternalServerError, fmt.Sprintf("strconv error: %s", err)})
		return nil, false
	}

	todo, err := s.TodoService.FindToDoByID(r.Context(), todoId)

	if err != nil {
		ReturnError(w, HttpError{http.StatusInternalServerError, fmt.Sprintf("db error: %s", err)})
		return nil, false
	}

	if todo == nil {
		// FindToDoById returns nil, nil if there is no matching id and no other error
		ReturnError(w, HttpError{http.StatusNotFound, "not found"})
		return nil, false
	}

	// Basic security check
	if todo.UserId != user.ID {
		ReturnError(w, HttpError{http.StatusForbidden, "not yours"})
		return nil, false
	}

	return todo, true
}
