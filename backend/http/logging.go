package http

import (
	"net/http"
	"time"
)

// Wrap the repsonse writer so we can get other info like the status code
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
	start      time.Time
}

func (rw *LoggingResponseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}

func (s *Server) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()

		rw := &LoggingResponseWriter{
			ResponseWriter: w,
			start:          now,
		}

		// Delegate to next HTTP handler.
		next.ServeHTTP(rw, r)

		duration := time.Since(rw.start)
		s.Logger.Printf("%s %s %d %dms", r.Method, r.URL, rw.statusCode, duration.Milliseconds())
	})
}
