package controllers

import (
	"NetGo/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

type Comment struct {
	Id   string    `json:"id"`
	Body string `json:"body"`
}

type Comments []Comment


func CommentsController(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodGet:
		param := utils.ExtractPathParam(request.URL.Path, "comments")
		writer.WriteHeader(http.StatusOK)

		if param != "" {
			response := Comment{Id: param, Body: "This is a comment"}
			json.NewEncoder(writer).Encode(response)
		} else {
			response := Comments{
				Comment{Id: "44", Body: "This is a comment"},
				Comment{Id: "56", Body: "This is another comment"},
			}
			json.NewEncoder(writer).Encode(response)

		}





	case http.MethodPost:
		// fmt.Println("[POST]", request.URL.Path)
	case http.MethodPut:
		// fmt.Println("[PUT]", request.URL.Path)
	case http.MethodDelete:
		// fmt.Println("[DELETE]", request.URL.Path)
	default:
		fmt.Println("Method not allowed")
		writer.WriteHeader(http.StatusMethodNotAllowed)
		writer.Write([]byte("Method not allowed"))
	}
}