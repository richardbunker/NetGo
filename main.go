package main

import (
	"NetGo/http"
	"NetGo/src/models/posts"
)

func main() {
	router := http.NetGo()
	router.RequireAuth()
	router.GET("/posts/:postId", posts.Show)
	router.GET("/posts", posts.Index)
	http.Start(3000, router)
}
