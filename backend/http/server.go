package http

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// All our HTTP server goodness
type Server struct {
	Host string
	Port int

	server *http.Server
	router *mux.Router
}

func NewServer() *Server {
	s := &Server{
		server: &http.Server{},
		router: mux.NewRouter(),
	}

	// No authenticated requests
	noAuthRouter := s.router.PathPrefix("/").Subrouter()
	noAuthRouter.HandleFunc("/uptest", s.handleUptest).Methods("GET")

	// use our mux
	s.server.Handler = s.router

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
