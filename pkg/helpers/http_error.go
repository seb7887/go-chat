package helpers

import (
	"net/http"
)

type HttpErrorResponse struct {
	Error string `json:"error"`
}

func HandleError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	RespondJSON(w, HttpErrorResponse{Error: message})
}
