```
    _   __     __  ______
   / | / /__  / /_/ ____/___
  /  |/ / _ \/ __/ / __/ __ \
 / /|  /  __/ /_/ /_/ / /_/ /
/_/ |_/\___/\__/\____/\____/

```

# NetGo

NetGo is a small and lightweight RESTful API application written in pure Go. It provides a simple way to create RESTful APIs. Features include:

- Route registration
- Request and response handling
- Middleware support
- JWT authentication
- DynamoDB integration
- User registration

## Route Registration

After instantiation, you may use the `Get`, `Post`, `Put`, and `Delete` methods to register routes.

```go
api := RestApi()

// Register the routes
api.Get("/users/:userId", ShowUser)
api.Put("/users/:userId", UpdateUser)
```

### Grouping Routes

You can group routes by using the `Group` method. This is useful for applying middleware to a group of routes or for clearer route organisation.

```go
api := RestApi()
api.Group("/users", []Middleware{}, func() {
	api.Get("/", ListUsers)
	api.Get("/:userId", ShowUser)
	api.Put("/:userId", UpdateUser)
})
```

The `ShowUser` and `UpdateUser` functions are the handlers for the routes. You can define them as follows:

## Route Handlers

Route handlers are functions that take a `RestApiRequest` as an argument and return a `RestApiResponse`. The `RestApiRequest` contains the request data, such as the path parameters, query parameters, and request body. The `RestApiResponse` contains the response data, such as the response body and status code.

```go
// A simple handler to show a user
func ShowUser(request RestApiRequest) RestApiResponse {
	userId := request.PathParams["userId"]
	return RestApiResponse{
		Body: map[string]interface{}{
			"id":   userId,
			"name": "Mickey Mouse",
		},
		StatusCode: 200,
	}
}
```

## Middleware

NetGo provides middleware support to allow you to proccess the incoming request before it reaches the route's handler function. You can use the `UseMiddleware` method to apply global middleware to you routes.

```go
api := RestApi()
api.UseMiddleware(LoggerMiddleware)
```

You can also apply middleware to a group of routes by passing the middleware to the `Group` method.

```go
api := RestApi()
api.Group("/users", []Middleware{Authenticate, RateLimit, Logger}, func() {
	api.Get("/", ListUsers)
	api.Get("/:userId", ShowUser)
	api.Put("/:userId", UpdateUser)
})
```

The middleware functions are processed in the order they are applied.

## JWT Authentication

NetGo provides JWT authentication support. You can use the `Authenticate` middleware to protect your routes.

```go
api := RestApi()
api.UseMiddleware(Authenticate)
```

## Deployment

To deploy the application, you can use the following command:

```bash
$ ./deploy.sh
```

Then you must upload the `deployment.zip` file to AWS Lambda and ensure your lambda function has the needed environment variables and DynamoDB permissions.

## Local development

To run the application locally, you can use the following command:

```bash
$ go run ./development/app.go
```

## Testing

To run the tests, you can use the following command:

```bash
$ ./test.sh
```
