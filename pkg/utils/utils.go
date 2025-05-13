package utils

import (
	"log"
	"net/http"
)

// HandleError logs an error message and sends an HTTP response with the error status.
func HandleError(w http.ResponseWriter, err error, status int) {
	log.Printf("Error: %v", err)
	http.Error(w, err.Error(), status)
}
