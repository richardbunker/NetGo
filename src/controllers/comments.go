package controllers

import (
	"NetGo/src/types"
	"NetGo/src/utils"
	"net/http"
)

type Comment struct {
	Id   int    `json:"id"`
	Body string `json:"body"`
}

type Comments []Comment

func CommentsController(method string, path string, session utils.Session) types.NetGoResponse  {
	param := utils.ExtractPathParam(path, "comments")
	var commentId int
	if param != "" {
		commentId = utils.ConvertPathParamToInt(param)
		if commentId == 0 {
			return types.NetGoResponse{Err: true, StatusCode: http.StatusBadRequest, Body: types.NetGoGenericResponse{Message: "Comment id must be an integer"}}
		}
	}
	switch method {
	case http.MethodGet:
		if param == "" {
			return indexComments()
		} else {
			return showComment(commentId)
		}
	case http.MethodPost:
		return createComment(commentId)
	case http.MethodPut:
		return updateComment(commentId)
	case http.MethodDelete:
		return deleteComment(commentId)
	default:
		return types.NetGoResponse{Err: true, StatusCode: http.StatusMethodNotAllowed, Body: types.NetGoGenericResponse{Message: "Method not allowed"}}
	}
}

func indexComments() types.NetGoResponse {
	return types.NetGoResponse{Err: false, StatusCode: http.StatusOK, Body: Comments{
		Comment{Id: 44, Body: "This is a comment"},
		Comment{Id: 56, Body: "This is another comment"},
	}}
}

func showComment(commentId int) types.NetGoResponse {
	return types.NetGoResponse{Err: false, StatusCode: http.StatusOK, Body: Comment{Id: commentId, Body: "This is a show comment"}}
}

func createComment(commentId int) types.NetGoResponse {
	if commentId > 0 {
		return types.NetGoResponse{Err: true, StatusCode: http.StatusBadRequest, Body: types.NetGoGenericResponse{Message: "Comment id not allowed"}}
	} else {
		return types.NetGoResponse{Err: false, StatusCode: http.StatusOK, Body: Comment{Id: 77, Body: "This is a newly created comment"}}
	}
}

func updateComment(commentId int) types.NetGoResponse {
	if commentId == 0 {
		return types.NetGoResponse{Err: true, StatusCode: http.StatusBadRequest, Body: types.NetGoGenericResponse{Message: "Comment id required"}}
		} else {
		return types.NetGoResponse{Err: false, StatusCode: http.StatusOK, Body: Comment{Id: commentId, Body: "This is a newly updated comment"}}
	}
}

func deleteComment(commentId int) types.NetGoResponse {
	if commentId == 0 {
		return types.NetGoResponse{Err: true, StatusCode: http.StatusBadRequest, Body: types.NetGoGenericResponse{Message: "Comment id required"}}
	} else {
		return types.NetGoResponse{Err: false, StatusCode: http.StatusOK, Body: types.NetGoGenericResponse{Message: "Comment deleted"}}
	}
}