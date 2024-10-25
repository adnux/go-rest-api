package utils

import (
	"encoding/json"
	"net/http"
)

func CreateHttpErrorMessage(w http.ResponseWriter, statusCode int, err error, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	errorResponse := map[string]string{
		"error":   err.Error(),
		"message": message,
	}
	json.NewEncoder(w).Encode(errorResponse)
}
