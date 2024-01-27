package controllers

import (
	"fmt"
	"net/http"
)

func PostsController(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		fmt.Println("[GET]", request.URL.Path)
	case http.MethodPost:
		fmt.Println("[POST]", request.URL.Path)
	case http.MethodPut:
		fmt.Println("[PUT]", request.URL.Path)
	case http.MethodDelete:
		fmt.Println("[DELETE]", request.URL.Path)
	default:
		fmt.Println("Method not allowed")
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed"))
	}
}
