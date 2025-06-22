package http

import (
	"net/http"
)

// Wrap the repsonse writer so we can get other info like the status code
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *LoggingResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (s *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &LoggingResponseWriter{
			ResponseWriter: w,
		}

		// Delegate to next HTTP handler.
		next.ServeHTTP(rw, r)

		s.Logger.Printf("%s %d", r.URL, rw.statusCode)
	})
}
