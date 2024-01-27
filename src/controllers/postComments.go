package controllers

import (
	"NetGo/src/types"
	"NetGo/src/utils"
	"net/http"
)

type PostWithComments struct {
	Id   int    `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Comments []Comment `json:"comments"`
}

func PostWithCommentsController(method string, path string) types.NetGoResponse  {
	param := utils.ExtractPathParam(path, "posts")
	if param == "" {
		return types.NetGoResponse{Err: true, StatusCode: http.StatusBadRequest, Body: types.NetGoGenericResponse{Message: "Post id required"}}
	}
	postId := utils.ConvertPathParamToInt(param)
	if postId == 0 {
		return types.NetGoResponse{Err: true, StatusCode: http.StatusBadRequest, Body: types.NetGoGenericResponse{Message: "Post id must be an integer"}}
	}
	switch method {
	case http.MethodGet:
		return showPostWithComments(postId)
	default:
		return types.NetGoResponse{Err: true, StatusCode: http.StatusMethodNotAllowed, Body: types.NetGoGenericResponse{Message: "Method not allowed"}}
	}
}

func showPostWithComments(postId int) types.NetGoResponse {
	return types.NetGoResponse{Err: false, StatusCode: http.StatusOK, Body: PostWithComments{Id: postId, Title: "Sample Title", Body: "This is a show Post", Comments: Comments{
		Comment{Id: 44, Body: "This is a comment"},
		Comment{Id: 56, Body: "This is another comment"},
		Comment{Id: 77, Body: "This is a newly created comment"},
	}}}
}