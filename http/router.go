package http

import (
	"NetGo/src/console"
	"NetGo/src/types"
	"NetGo/src/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

type Router struct {
	requireAuth bool
	session *utils.Session
	methods types.Methods
	params map[string]string
}

func NetGo() *Router {
	return &Router{
		requireAuth: false,
		session: nil,		
		methods: make(types.Methods),
		params: make(map[string]string),
	}
}

func (r *Router) RequireAuth() {
	r.requireAuth = true
}

func (r *Router) GET(path string, handle func(w http.ResponseWriter, r *http.Request)) {
	r.Handle(http.MethodGet, path, handle)
}

func (r *Router) Handle(method string, path string, handle func(w http.ResponseWriter, r *http.Request)) {
	if r.methods[method] == nil {
		r.methods[method] = make(types.Routes)
	}
	r.methods[method][path] = handle
}


func (router *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")
	if router.requireAuth {
		jwt := r.Header.Get("Authorization")
		if jwt == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(types.NetGoGenericResponse{Message: "Unauthorized"})
			return
		}
		router.session = utils.ParseJWT(jwt)
		if router.session == nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(types.NetGoGenericResponse{Message: "Unauthorized"})
			return
		}
	}
	incomingPath := r.URL.Path
	for route := range router.methods[r.Method] {
		if match(incomingPath, route) {
			handler := router.methods[r.Method][route]
			handler(w, r)
			console.LogRequest(r.Method, incomingPath, router.session)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(types.NotFoundResponse{Message: "Not Found"})
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