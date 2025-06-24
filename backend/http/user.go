package http

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (s *Server) registerUserRoutes(r *mux.Router) {
	// Listing of all todos for a user.
	r.HandleFunc("/api/1.0/user", s.handleUserIndex).Methods("GET")
}

func (s *Server) handleUserIndex(w http.ResponseWriter, r *http.Request) {
	// The authn middleware will ensure we have alogged in user here so shouldn't fail
	user, ok := s.getUser(w, r)
	if !ok {
		return
	}

	ReturnJson(w, http.StatusOK, user)
}
