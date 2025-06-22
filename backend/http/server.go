package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
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
}

func NewServer(sessionConfig SessionConfig) *Server {
	s := &Server{
		server: &http.Server{},
		router: mux.NewRouter(),
	}

	// No authenticated requests
	noAuthRouter := s.router.PathPrefix("/").Subrouter()
	noAuthRouter.HandleFunc("/uptest", s.handleUptest).Methods("GET")

	// Register all other routes
	s.registerTodoRoutes(s.router)

	// Create the session manager (the app will set the store later once it knows what DB)
	s.SessionManager = scs.New()
	s.SessionManager.Lifetime = sessionConfig.Lifetime
	s.SessionManager.IdleTimeout = sessionConfig.IdleTimeout
	s.SessionManager.Cookie.Secure = sessionConfig.Secure
	s.SessionManager.Cookie.Name = sessionConfig.Name
	s.SessionManager.Cookie.SameSite = sessionConfig.SameSite

	// Wrap the gorilla mux in session and logging so that logger is the parent
	// followed next by session, then it's whatever is set via gorilla mux
	//
	// We use scs rather than roll our own because hopefully a mature session manager has
	// had the security bugs fixed.
	s.server.Handler = s.logRequest(s.SessionManager.LoadAndSave(s.router))

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

func ReturnJson(w http.ResponseWriter, code int, content any) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(content)
}
