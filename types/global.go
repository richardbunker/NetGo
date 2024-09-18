package types

type RestApiRequest struct {
	Headers    map[string][]string
	Cookies    []string
	Method     string
	Path       string
	PathParams map[string]string
	Query      map[string]string
	Body       map[string]interface{}
}

type RestApiResponse struct {
	Body       interface{}
	StatusCode int
}

type Handler func(request RestApiRequest) RestApiResponse

type MiddlewareReason struct {
	StatusCode int
	Message    string
}

type Middleware func(request RestApiRequest) (error, *MiddlewareReason)

type RouteOptions struct {
	Middleware []Middleware
	Handler    Handler
}

type Routes map[string]RouteOptions

type Method string

const (
	GET    Method = "GET"
	POST   Method = "POST"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
)

type User struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
