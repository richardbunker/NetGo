package postsController

import (
	"fmt"
	"net/http"
)

func Handler(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		Index(writer, request)
	case http.MethodPost:
		Create(writer, request)
	default:
		fmt.Println("Method not allowed")
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed"))
	}
}

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Index Post")
	writer.Write([]byte("Index Post"))
}

func Create(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Create Post")
	writer.Write([]byte("Create Post"))
}