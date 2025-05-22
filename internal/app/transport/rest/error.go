package rest

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Description string `json:"error"`
}

func WriteErrorResponse(w http.ResponseWriter, statusCode int, description string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(ErrorResponse{Description: description})
}
