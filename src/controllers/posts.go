package controllers

import (
	"NetGo/src/types"
	"NetGo/src/utils"
	"net/http"
)

type Post struct {
	Id   int    `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}

type Posts []Post

func PostsController(method string, path string) types.NetGoResponse  {
	param := utils.ExtractPathParam(path, "posts")
	var postId int
	if param != "" {
		postId = utils.ConvertPathParamToInt(param)
		if postId == 0 {
			return types.NetGoResponse{Err: true, StatusCode: http.StatusBadRequest, Body: types.NetGoGenericResponse{Message: "Post id must be an integer"}}
		}
	}
	switch method {
	case http.MethodGet:
		if param == "" {
			return indexPosts()
		} else {
			return showPost(postId)
		}
	case http.MethodPost:
		return createPost(postId)
	case http.MethodPut:
		return updatePost(postId)
	case http.MethodDelete:
		return deletePost(postId)
	default:
		return types.NetGoResponse{Err: true, StatusCode: http.StatusMethodNotAllowed, Body: types.NetGoGenericResponse{Message: "Method not allowed"}}
	}
}

func indexPosts() types.NetGoResponse {
	return types.NetGoResponse{Err: false, StatusCode: http.StatusOK, Body: Posts{
		Post{Id: 44, Title: "Sample Title", Body: "This is a Post"},
		Post{Id: 56, Title: "Sample Title", Body: "This is another Post"},
	}}
}

func showPost(postId int) types.NetGoResponse {
	return types.NetGoResponse{Err: false, StatusCode: http.StatusOK, Body: Post{Id: postId, Title: "Sample Title", Body: "This is a show Post"}}
}

func createPost(postId int) types.NetGoResponse {
	if postId > 0 {
		return types.NetGoResponse{Err: true, StatusCode: http.StatusBadRequest, Body: types.NetGoGenericResponse{Message: "Post id not allowed"}}
	} else {
		return types.NetGoResponse{Err: false, StatusCode: http.StatusOK, Body: Post{Id: 77, Title: "Sample Title", Body: "This is a newly created Post"}}
	}
}

func updatePost(postId int) types.NetGoResponse {
	if postId == 0 {
		return types.NetGoResponse{Err: true, StatusCode: http.StatusBadRequest, Body: types.NetGoGenericResponse{Message: "Post id required"}}
		} else {
		return types.NetGoResponse{Err: false, StatusCode: http.StatusOK, Body: Post{Id: postId, Title: "Sample Title", Body: "This is a newly updated Post"}}
	}
}

func deletePost(postId int) types.NetGoResponse {
	if postId == 0 {
		return types.NetGoResponse{Err: true, StatusCode: http.StatusBadRequest, Body: types.NetGoGenericResponse{Message: "Post id required"}}
	} else {
		return types.NetGoResponse{Err: false, StatusCode: http.StatusOK, Body: types.NetGoGenericResponse{Message: "Post deleted"}}
	}
}