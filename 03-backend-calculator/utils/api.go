package utils

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	WriteJSON(w, status, ErrorResponse{Error: message})
}

func ValidateJsonBody(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, io.EOF):
		WriteError(w, http.StatusBadRequest, "Request body is empty")
	case errors.Is(err, io.ErrUnexpectedEOF):
		WriteError(w, http.StatusBadRequest, "Request body contains invalid JSON")
	default:
		WriteError(w, http.StatusBadRequest, "Invalid JSON format")
	}
	return
}
