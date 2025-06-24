package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rtbaker/GoToDo/password"
)

// Login/logout etc

func (s *Server) registerAuthnRoutes(r *mux.Router) {
	// Listing of all todos for a user.
	r.HandleFunc("/api/1.0/login", s.login).Methods("POST")
	r.HandleFunc("/api/1.0/logout", s.logout).Methods("POST")
}

type LoginDTO struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	var loginInfo LoginDTO

	if err := json.NewDecoder(r.Body).Decode(&loginInfo); err != nil {
		ReturnError(w, HttpError{http.StatusBadRequest, "invalid JSON body"})
		return
	}

	if err := s.validator.Struct(loginInfo); err != nil {
		ReturnError(w, HttpError{http.StatusBadRequest, fmt.Sprintf("validation error: %s", err)})
		return
	}

	user, err := s.UserService.FindUserByEmail(r.Context(), loginInfo.Email)

	if err != nil {
		ReturnError(w, HttpError{http.StatusBadRequest, "email/password does not match"})
		return
	}

	if !password.VerifyPassword(loginInfo.Password, user.Password) {
		ReturnError(w, HttpError{http.StatusBadRequest, "email/password does not match"})
		return
	}

	// Got this far, then login must be good, sort out the session and send the user info back
	err = s.SessionManager.RenewToken(r.Context())
	if err != nil {
		ReturnError(w, HttpError{http.StatusInternalServerError, fmt.Sprintf("sessions error: %s", err)})
		return
	}

	s.SessionManager.Put(r.Context(), "user", user)

	ReturnJson(w, http.StatusOK, user)
}

func (s *Server) logout(w http.ResponseWriter, r *http.Request) {
	// We just destroy the session regardless of whether anyone is logged in
	s.SessionManager.Destroy(r.Context())

	content := map[string]string{"status": "ok"}
	ReturnJson(w, http.StatusOK, content)
}
