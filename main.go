package main

import (
	. "NetGo/app"
	. "NetGo/lib"
	. "NetGo/types"
	"fmt"
	"net/http"
	"strconv"
)

// The port the server will run on
var port int = 5050

func main() {
	// Create a new API application
	app := RestApi()
	app.Post("/posts/:postId", RouteOptions{
		Handler: func(request RestApiRequest) RestApiResponse {
			return RestApiResponse{
				Body: map[string]interface{}{
					"postId": request.PathParams["postId"],
				},
				StatusCode: 200,
			}
		},
	})
	StartUpMessage(strconv.Itoa(port))
	if err := http.ListenAndServe(":"+strconv.Itoa(port), app); err != nil {
		fmt.Printf("Server failed to start: %s\n", err)
	}
}
