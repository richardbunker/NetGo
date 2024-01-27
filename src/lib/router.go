package lib

import (
	"NetGo/src/console"
	"NetGo/src/types"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

// The main router function to determine which controller to call
func Router(w http.ResponseWriter, r *http.Request, routeList types.RouteList ) {
	console.LogRequest(r.Method, r.URL.Path)

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Router forwards requests to each controller
	var handler = selectControllerToHandleRequest(r.URL.Path, routeList) 
	if handler != nil {
		handler(w, r)
	} else {
		w.WriteHeader(http.StatusNotFound)
		response := types.NotFoundResponse{Message: "Not Found"}
		json.NewEncoder(w).Encode(response)
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

// A function to loop through an array of routes and call the appropriate controller
func selectControllerToHandleRequest(path string, routeList types.RouteList) func(writer http.ResponseWriter, request *http.Request) {
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