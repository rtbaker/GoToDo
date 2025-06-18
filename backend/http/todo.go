package http

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) registerTodoRoutes(r *mux.Router) {
	// Listing of all todos for a user.
	r.HandleFunc("/api/1.0/todos", s.handleTodoIndex).Methods("GET")
}

func (s *Server) handleTodoIndex(w http.ResponseWriter, r *http.Request) {
	todos, err := s.TodoService.FindByUser(r.Context(), 13)

	if err != nil {
		httpErr := HttpError{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("%s", err),
		}

		ReturnError(w, httpErr)
	}

	ReturnJson(w, http.StatusOK, todos)
}
