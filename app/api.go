package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	. "rest_api/lib"
	. "rest_api/types"
	"strings"
)

type Api struct {
	methods          map[Method]Routes
	basePath         string
	globalMiddleware []Middleware
}

// Create a new API application
func RestApi() *Api {
	return &Api{
		methods:          make(map[Method]Routes),
		basePath:         "",
		globalMiddleware: []Middleware{},
	}
}

func (api *Api) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	request := RestApiAdaptor(r)
	response := api.HandleRequest(request)
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response.Body)
}

// Register a handler for a GET request
func (api *Api) Get(pathToMatch string, routeOptions RouteOptions) {
	api.RegisterRoute(GET, pathToMatch, routeOptions)
}

// Register a handler for a POST request
func (api *Api) Post(pathToMatch string, routeOptions RouteOptions) {
	api.RegisterRoute(POST, pathToMatch, routeOptions)
}

// Register a handler for a PUT request
func (api *Api) Put(pathToMatch string, routeOptions RouteOptions) {
	api.RegisterRoute(PUT, pathToMatch, routeOptions)
}

// Register a handler for a DELETE request
func (api *Api) Delete(pathToMatch string, routeOptions RouteOptions) {
	api.RegisterRoute(DELETE, pathToMatch, routeOptions)
}

// Register a group of routes
func (api *Api) Group(basePath string, middlewares []Middleware, registerRoutes func()) {
	originalBasePath := api.basePath
	originalMiddlewares := api.globalMiddleware

	api.basePath = originalBasePath + basePath
	api.globalMiddleware = append(originalMiddlewares, middlewares...)

	registerRoutes()

	api.basePath = originalBasePath
	api.globalMiddleware = originalMiddlewares
}

func validateRegisteredRoute(pathToMatch string) error {
	invalidCharsMap := map[string]bool{
		" ": true,
		"$": true,
		"^": true,
		"[": true,
		"]": true,
		"{": true,
		"}": true,
		"(": true,
		")": true,
	}
	for _, char := range pathToMatch {
		_, ok := invalidCharsMap[string(char)]
		if ok {
			return fmt.Errorf("Invalid character in path: %s", string(char))
		}
	}
	return nil
}

// Register a handler for a request
func (api *Api) RegisterRoute(method Method, pathToMatch string, routeOptions RouteOptions) {
	err := validateRegisteredRoute(pathToMatch)
	if err != nil {
		return
	}
	if api.methods[method] == nil {
		api.methods[method] = make(Routes)
	}
	routeOptions.Middleware = append(api.globalMiddleware, routeOptions.Middleware...)
	api.methods[method][api.basePath+pathToMatch] = routeOptions
}

// Handle a request
func (api *Api) HandleRequest(request RestApiRequest) RestApiResponse {
	err := validateMethodHasRoutes(Method(request.Method), api.methods)
	if err != nil {
		return ApiErrorResponse(405, "Method not allowed")
	}
	var routeOptions RouteOptions
	for route := range api.methods[Method(request.Method)] {
		if match(request.Path, route) {
			routeOptions = api.methods[Method(request.Method)][route]
			request.PathParams = extractPathParams(request.Path, route)
			break
		}
	}
	if routeOptions.Handler == nil {
		return ApiErrorResponse(404, "Route not found")
	}
	for _, middleware := range routeOptions.Middleware {
		error, reason := middleware(request)
		if error != nil {
			return ApiErrorResponse(reason.StatusCode, reason.Message)
		}
	}
	LogRequest(request)
	return routeOptions.Handler(request)
}

func validateMethodHasRoutes(requestedMethod Method, methods map[Method]Routes) error {
	if methods[Method(requestedMethod)] == nil {
		return fmt.Errorf("Method %s has no registered routes", requestedMethod)
	}
	return nil
}

// Extract path parameters from the requested path
func extractPathParams(requestedPath string, registeredPath string) map[string]string {
	params := make(map[string]string)
	requestedPathParts := strings.Split(requestedPath, "/")
	registeredPathParts := strings.Split(registeredPath, "/")
	for i, part := range registeredPathParts {
		if strings.HasPrefix(part, ":") {
			params[strings.TrimPrefix(part, ":")] = requestedPathParts[i]
		}
	}
	return params
}

// Check if the requested path matches the registered route
func match(path string, route string) bool {
	pattern := regexp.MustCompile(`:\w+`).ReplaceAllString(route, "([^/]+)")
	regex, _ := regexp.Compile("^" + pattern + "$")
	return regex.MatchString(path)
}

// A helper function to create an error response
func ApiErrorResponse(statusCode int, message string) RestApiResponse {
	return RestApiResponse{
		Body: map[string]interface{}{
			"error": message,
		},
		StatusCode: statusCode,
	}
}

func (api *Api) Middleware(middleware []Middleware) {
	for _, m := range middleware {
		api.globalMiddleware = append(api.globalMiddleware, m)
	}
}
