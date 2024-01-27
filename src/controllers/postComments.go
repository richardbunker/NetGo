package controllers

import (
	"fmt"
	"net/http"
)

func PostCommentsController(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		fmt.Println("[GET]", request.URL.Path)
	default:
		fmt.Println("Method not allowed")
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed"))
	}
}
