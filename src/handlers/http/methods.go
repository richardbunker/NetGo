package ng_method_handlers

import (
	"encoding/json"
	"net/http"
)

type Json struct {
	Message string `json:"message"`
}

func HandleGet(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Return JSON
	json.NewEncoder(w).Encode(Json{Message: "Hello World!"})
}