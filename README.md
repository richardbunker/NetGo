```
    _   __     __  ______
   / | / /__  / /_/ ____/___
  /  |/ / _ \/ __/ / __/ __ \
 / /|  /  __/ /_/ /_/ / /_/ /
/_/ |_/\___/\__/\____/\____/

```

# NetGo

NetGo is an AWS 'Lambda-first' RESTful API application written in pure Go. It provides a simple way to create RESTful APIs with minimal dependencies. It is designed to be lightweight and easy to use. It is built for the AWS cloud and tightly integrates with AWS services such as DynamoDB, API Gateway, Certificate Manager, and Lambda.

## ✨ Features

-   Route registration
-   Request and response handling
-   Middleware support
-   JWT authentication
-   DynamoDB integration
-   User registration
-   Full Terraform deployment support
-   Local development support
-   Unit tests [WIP]

## 🚀 Quick Start

To get started with NetGo, clone the repository:

```bash
$ git clone git@github.com:richardbunker/NetGo.git
```

Then, `cd` into NetGo and run the following command:

```bash
$ go run ./development/app.go
```

This will start a local server on port 8080. You can then access the server at `http://localhost:8080`.

## 📦 Deployment

NetGo is designed to be deployed to AWS via Terraform. To deploy the application, you can use the following command:

```bash
$ cd ./terraform
$ terraform init

# ./deploy.sh <AWS_PROFILE> <TERRAFORM_ACTION>
$ ./deploy.sh my-aws-profile plan
```

Running `./deploy.sh` will build the Go application, create a `deployment.zip`, and deploy the application to AWS using Terraform. The `AWS_PROFILE` is the AWS CLI profile to use, and the `TERRAFORM_ACTION` is the Terraform action to perform, such as `plan` or `apply`.

##### Note 💡: The `deploy.sh` script assumes that you have the AWS CLI and Terraform installed and setup on your machine.

### Terraform Variables

You will need to set the following Terraform variables in a new `vars.tfvars` file:

```bash
$ vim vars.tfvars
```

An example `vars.tfvars.example` file is provided in the `terraform` directory:

```hcl
service_name         = "your-app-name-here"
cost_center          = "your-app-name-here-cc"
lambda_function_name = "your-app-name-here-lambda"
api_name             = "your-app-name-here-api"
dynamodb_table_name  = "your-app-name-here-table"
dynamodb_global_secondary_index_name = "your-app-name-here-gsi"
lambda_dynamodb_policy_name = "your-app-name-here-dynamodb-policy"
api_domain_name = "your-app-name-here.your-domain.com"
domain_name = "your-domain.com"
email_from="your@email.com"
email_password="your-password"
email_smtp_host="your.host"
email_smtp_port=587
```

### AWS Route53 Hosted Zone

NetGo requires an existing AWS Route53 Hosted Zone along with a registered domain. This will need to be done manually before running the Terraform deployment.

To do this see the following AWS documentation: [Creating a Public Hosted Zone](https://docs.aws.amazon.com/Route53/latest/DeveloperGuide/CreatingHostedZone.html)

## 🛣️ Route Registration

The `./app/bootstrap.go` is where you may register your routes. The `RestApi` function returns a new instance of the `RestApi` struct. After instantiation, you may use the `Get`, `Post`, `Put`, and `Delete` methods to register routes.

```go
api := RestApi()
// Register the routes
api.Get("/posts/:postId", ShowPost)
```

### Grouping Routes

You can group routes by using the `Group` method. This is useful for applying middleware to a group of routes or for clearer route organisation.

```go
api := RestApi()
// Register the Post CRUD routes
api.Group("/posts", []Middleware{}, func() {
    api.Get("/", ListPosts)
    api.Get("/:postId", ShowPost)
    api.Post("/", CreatePost)
    api.Put("/:postId", UpdatePost)
    api.Delete("/:postId", DeletePost)
})
```

## 🔧 Route Handlers

Route handlers are functions that take a `RestApiRequest` as an argument and return a `RestApiResponse`. The `RestApiRequest` contains the request data, such as the path parameters, query parameters, and request body. The `RestApiResponse` contains the response data, such as the response body and status code.

```go
// A simple handler to show a user
func ShowUser(request RestApiRequest) RestApiResponse {
	userId := request.PathParams["userId"]
	user := ... // Fetch the user from the database
	return RestApiResponse{
		Body: user,
		StatusCode: 200,
	}
}
```

## 🧩 Middleware

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

##### Note 💡: The middleware functions are processed in the order they are applied.

## 🖊️ User Registration

NetGo provides user registration with magic login link support. You can use the following endpoint to register a new user:

```bash
POST /auth/register

{
    "email": "your@email.com",
    "name": "Your Name",
}
```

Once you successfully register a user, you may send a post request to the following endpoint to receive a magic login link:

```bash
POST /auth/email-magic-link

{
    "email": "your@email.com"
}
```

A magic login link will be sent to the user's email address. The user can then click the link to log in. The link essentially contains an opaque token that is used to authenticate the user and to generate a JWT token. To obtain the JWT token, you can send a post request to the following endpoint:

```bash
POST /auth/login

{
    "token": "pO9EMekMwnwTCXtdZl74Iskm9BHa6YaMInVDSlA4d8duCJW3"
}
```

Expected response:

```bash
{
    "jwt": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
}
```

## 🔐 JWT Authentication

NetGo provides JWT authentication support out of the box. You can use the `Authenticate` middleware to protect your routes.

```go
api := RestApi()
api.UseMiddleware(Authenticate)
```

The `Authenticate` middleware will check the `Authorization` header for a valid JWT token. If the token is valid, the request will be passed to the route's handler function. If the token is invalid, the middleware will return a `401 Unauthorized` response.

```bash
// Example request with JWT token
curl -X GET http://localhost:8080/users/1 -H "Authorization eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"
```

## 💻 Local development

To run the application locally, you can use the following command:

```bash
$ go run ./development/app.go
```

## 🧪 Testing

To run the tests, you can use the following command:

```bash
$ ./test.sh
```
