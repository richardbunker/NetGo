package types

type RouteList map[string]func(method string, path string) NetGoResponse
// type RouteList map[string]func(w http.ResponseWriter, r *http.Request)

type NotFoundResponse struct {
	Message string `json:"message"`
}

type NetGoGenericResponse struct {
	Message string `json:"message"`
}

type NetGoResponse struct {
	Err 	bool 		`json:"err"`
	StatusCode int 		`json:"statusCode"`
	Body 	interface{} `json:"body"`
}