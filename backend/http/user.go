package http

import (
	"net/http"

	"github.com/gorilla/mux"
	gotodo "github.com/rtbaker/GoToDo/Model"
)

func (s *Server) registerUserRoutes(r *mux.Router) {
	// Listing of all todos for a user.
	r.HandleFunc("/api/1.0/user", s.handleUserIndex).Methods("GET")
}

func (s *Server) handleUserIndex(w http.ResponseWriter, r *http.Request) {
	// The authn middleware will ensur ewe have alogge din user here so shouldn't fail
	user, ok := s.SessionManager.Get(r.Context(), "user").(gotodo.User)
	if !ok {
		httpErr := HttpError{
			Code:    http.StatusInternalServerError,
			Message: "cannot conver user value in session",
		}

		ReturnError(w, httpErr)
	}

	ReturnJson(w, http.StatusOK, user)
}
