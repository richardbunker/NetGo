```
    _   __     __  ______
   / | / /__  / /_/ ____/___
  /  |/ / _ \/ __/ / __/ __ \
 / /|  /  __/ /_/ /_/ / /_/ /
/_/ |_/\___/\__/\____/\____/

```

# NetGo

NetGo is an AWS 'Lambda-first' RESTful API application written in pure Go. It provides a simple way to create RESTful APIs with minimal dependencies. It is designed to be lightweight and easy to use. It is built for the AWS cloud and tightly integrates with AWS services such as DynamoDB, API Gateway, Certificate Manager, and Lambda.

## Features

- Route registration
- Request and response handling
- Middleware support
- JWT authentication
- DynamoDB integration
- User registration
- Full Terraform deployment support
- Local development support
- Unit tests [WIP]

## Deployment

NetGo is designed to be deployed to AWS via Terraform. To deploy the application, you can use the following command:

```bash
$ cd ./terraform
$ terraform init

# ./deploy.sh <AWS_PROFILE> <TERRAFORM_ACTION>
$ ./deploy.sh my-aws-profile plan
```

Running `./deploy.sh` will build the Go application, create a ZIP archive, and deploy the application to AWS using Terraform. The `AWS_PROFILE` is the AWS CLI profile to use, and the `TERRAFORM_ACTION` is the Terraform action to perform, such as `plan` or `apply`.

#### Note: The `deploy.sh` script assumes that you have the AWS CLI and Terraform installed and setup on your machine.

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
