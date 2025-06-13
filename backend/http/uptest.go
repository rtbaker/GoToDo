package http

import "net/http"

// The uptest route
func (s *Server) handleUptest(w http.ResponseWriter, r *http.Request) {
	content := map[string]string{"status": "ok"}
	ReturnJson(w, http.StatusOK, content)
}
