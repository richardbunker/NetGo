package posts

import (
	"encoding/json"
	"net/http"
)


func Index(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{"message": "Hello Posts!"})
}