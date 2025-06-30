package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	gotodo "github.com/rtbaker/GoToDo/Model"
)

// All our HTTP server goodness
type Server struct {
	Host string
	Port int

	server *http.Server
	router *mux.Router

	TodoService    gotodo.ToDoService
	UserService    gotodo.UserService
	SessionManager *scs.SessionManager

	Logger *log.Logger

	validator *validator.Validate
}

func NewServer(sessionConfig SessionConfig) *Server {
	s := &Server{
		server: &http.Server{},
		router: mux.NewRouter(),
	}

	// CORS setup, TODO: options in config file
	corsOrigins := handlers.AllowedOrigins([]string{"http://localhost:5173"})
	corsHeaders := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Content-Language", "Origin"})
	corsMethods := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "DELETE", "PUT"})
	corsCredentials := handlers.AllowCredentials()
	corsHandler := handlers.CORS(corsOrigins, corsHeaders, corsMethods, corsCredentials)

	// No authenticated requests
	noAuthRouter := s.router.PathPrefix("/").Subrouter()
	noAuthRouter.HandleFunc("/uptest", s.handleUptest).Methods("GET")
	s.registerAuthnRoutes(noAuthRouter)

	// Register all authenticated routes
	authRouter := s.router.PathPrefix("/").Subrouter()
	authRouter.Use(s.requireAuth)
	s.registerTodoRoutes(authRouter)
	s.registerUserRoutes(authRouter)

	// Create the session manager (the app will set the store later once it knows what DB)
	s.SessionManager = scs.New()
	s.SessionManager.Lifetime = sessionConfig.Lifetime
	s.SessionManager.IdleTimeout = sessionConfig.IdleTimeout
	s.SessionManager.Cookie.Secure = sessionConfig.Secure
	s.SessionManager.Cookie.Name = sessionConfig.Name
	s.SessionManager.Cookie.SameSite = sessionConfig.SameSite

	// Just the vanilla validator for now
	s.validator = validator.New()

	// Wrap the gorilla mux in session and logging so that logger is the parent
	// followed next by session, then it's whatever is set via gorilla mux
	//
	// We use scs rather than roll our own because hopefully a mature session manager has
	// had the security bugs fixed.
	s.server.Handler = s.logRequest(corsHandler(s.SessionManager.LoadAndSave(s.router)))

	return s
}

func (s *Server) Run() error {
	s.server.Addr = fmt.Sprintf("%s:%d", s.Host, s.Port)

	go s.server.ListenAndServe()

	return nil
}

func (s *Server) Close() error {
	return nil
}

// TODO: add a disable flag to a user (and do user refresh on each request to check for changes)
func (s *Server) requireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// the session will have a user object
		if !s.SessionManager.Exists(r.Context(), "user") {
			httpErr := HttpError{
				Code:    http.StatusUnauthorized,
				Message: "no authentication",
			}

			ReturnError(w, httpErr)

			return
		}

		// No further checks for now
		next.ServeHTTP(w, r)
	})
}

// Gets the user from the context/session and returns it, if no user available
// then return error. To be used by authenticated endpoints.
func (s *Server) getUser(w http.ResponseWriter, r *http.Request) (gotodo.User, bool) {
	user, ok := s.SessionManager.Get(r.Context(), "user").(gotodo.User)
	if !ok {
		httpErr := HttpError{
			Code:    http.StatusInternalServerError,
			Message: "cannot conver user value in session",
		}

		ReturnError(w, httpErr)
	}

	return user, true
}

func ReturnJson(w http.ResponseWriter, code int, content any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(content)
}
