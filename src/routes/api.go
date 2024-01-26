package routes

import (
	"NetGo/src/lib"
	"fmt"
	"net/http"
	"regexp"
)

var GET, POST = lib.GET, lib.POST

func Router(w http.ResponseWriter, r *http.Request) {
	// http.HandleFunc("/posts", postsController.Handler)
	path := r.URL.Path
	pattern :=  regexp.MustCompile(`:\w+`).ReplaceAllString("/posts/:id/comments", "([^/]+)")
	regex, err := regexp.Compile("^" + pattern + "$")
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Internal Server Error"}`))
	}
	if regex.MatchString(path) {
		fmt.Println("Matched")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Matched"))
	} 
	
	// http.HandleFunc("/posts/:id/comments", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Comments"))
	// })
	// POST("/posts", posts.Create)
	// Put("/posts/:id", posts.Update)
	// Delete("/posts/:id", posts.Delete)
}