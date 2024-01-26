package lib

import (
	"fmt"
	"net/http"
)

func GET(path string, handler func(http.ResponseWriter, *http.Request)) {
	fmt.Println("GET")
	http.HandleFunc(path, handler)
}

func POST(path string, handler func(http.ResponseWriter, *http.Request)) {
	fmt.Println("POST")	
	if http.MethodPost == "POST" {
		http.HandleFunc(path, handler)
	}
}