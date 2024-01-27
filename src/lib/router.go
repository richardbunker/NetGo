package lib

import (
	"NetGo/src/console"
	"NetGo/src/types"
	"NetGo/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

// The main router function to determine which controller to call
func Router(w http.ResponseWriter, r *http.Request, routeList types.RouteList ) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	
	// Validate the JWT
	jwt := r.Header.Get("Authorization")
	if jwt == "" {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(types.NetGoGenericResponse{Message: "Unauthorized"})
		return
	}
	session := utils.ParseJWT(jwt)
	if session == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(types.NetGoGenericResponse{Message: "Unauthorized"})
		return
	}
	console.LogRequest(r.Method, r.URL.Path, session)
	
	// Router forwards requests to each controller
	var controller = selectControllerToHandleRequest(r.URL.Path, routeList) 
	if controller != nil {
		response := controller(r.Method, r.URL.Path, *session)
		if response.Err {
			w.WriteHeader(response.StatusCode)
			json.NewEncoder(w).Encode(response.Body)
		} else {
			w.WriteHeader(response.StatusCode)
			json.NewEncoder(w).Encode(response.Body)
		}
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(types.NotFoundResponse{Message: "Not Found"})
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
func selectControllerToHandleRequest(path string, routeList types.RouteList) func(method string, path string, session utils.Session) types.NetGoResponse {
	// Loop through the routes and call the appropriate controller
	var matchedController func(method string, path string, session utils.Session) types.NetGoResponse
	for route, controller := range routeList {
		if match(path, route) {
			matchedController = controller
			break
		}
	}
	return matchedController
}