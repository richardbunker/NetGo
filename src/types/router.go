package types

import (
	"net/http"
)

type RouteList map[string]func(w http.ResponseWriter, r *http.Request)

type NotFoundResponse struct {
	Message string `json:"message"`
}