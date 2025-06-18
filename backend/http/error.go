package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HttpError struct {
	// Machine-readable error code.
	Code int

	// Human-readable error message.
	Message string
}

// Error implements the error interface. Not used by the application otherwise.
func (e HttpError) Error() string {
	return fmt.Sprintf("gotodo error: code=%d message=%s", e.Code, e.Message)
}

func ReturnError(w http.ResponseWriter, err HttpError) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(err.Code)
	json.NewEncoder(w).Encode(err)
}
