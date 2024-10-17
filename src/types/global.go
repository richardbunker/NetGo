package types

type NetGoRequest struct {
	Headers    map[string][]string
	Cookies    []string
	Method     string
	Path       string
	PathParams map[string]string
	Query      map[string]string
	Body       map[string]interface{}
}

type NetGoResponse struct {
	Body       interface{}
	StatusCode int
}

type Handler func(request NetGoRequest) NetGoResponse

type MiddlewareReason struct {
	StatusCode int
	Message    string
}

type Middleware func(request NetGoRequest) (error, *MiddlewareReason)

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

// User is a simplified struct to work with users from the database.
type User struct {
	Id    string
	Email string
	Name  string
}

// UserJson is the struct to help prepare the JSON response for users.
type UserJson struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// UserDynamoDBItem is the struct to help prepare the DynamoDB put item action for users.
type UserDynamoDBItem struct {
	PK    string
	SK    string
	GSIPK string
	GSISK string
	Type  string
	Name  string
}

// A simplified stuct to work with login tokens from the database.
type LoginToken struct {
	Id        string
	UserId    string
	ExpiresAt string
}

// LoginTokenDynamoDB is the struct to help prepare the DynamoDB put item action for login tokens.
type LoginTokenDynamoDBItem struct {
	PK        string
	SK        string
	GSIPK     string
	Type      string
	ExpiresAt string
}
