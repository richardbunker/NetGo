package routes

import (
	"NetGo/src/controllers"
	"fmt"
	"net/http"
	"regexp"
)


func Router(w http.ResponseWriter, r *http.Request) {
	// A map of routes and their controllers
	var routeList = map[string]func(writer http.ResponseWriter, request *http.Request){
		"/posts": controllers.PostsController,
		"/posts/:id": controllers.PostsController,
		"/posts/:id/comments": controllers.PostCommentsController,
	}

	// Router forwards requests to each controller
	var handler = selectControllerToHandleRequest(r.URL.Path, routeList) 
	if handler != nil {
		handler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Not Found"))
	}
}

// A function to loop through an array of routes and call the appropriate controller
func match(path string, route string) bool {
	pattern :=  regexp.MustCompile(`:\w+`).ReplaceAllString(route, "([^/]+)")
	regex, err := regexp.Compile("^" + pattern + "$")
	if err != nil {
		fmt.Println("Error in regex")	
		return false;
	}
	return regex.MatchString(path)
}


func selectControllerToHandleRequest(path string, routeList map[string]func(w http.ResponseWriter, r *http.Request)) func(writer http.ResponseWriter, request *http.Request) {
	// A map of routes and their controllers
	
	// Loop through the routes and call the appropriate controller
	var matchedController func(writer http.ResponseWriter, request *http.Request)
	for route, controller := range routeList {
		if match(path, route) {
			matchedController = controller
			break
		}
	}
	return matchedController
}