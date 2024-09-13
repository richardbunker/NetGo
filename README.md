```
    _   __     __  ______
   / | / /__  / /_/ ____/___
  /  |/ / _ \/ __/ / __/ __ \
 / /|  /  __/ /_/ /_/ / /_/ /
/_/ |_/\___/\__/\____/\____/

```

# NetGo

NetGo is a simple REST API framework for Go. It provides a simple way to create REST APIs and register routes.

## Route Registration

Register your routes in the `main.go` file. You can use the `Get`, `Post`, `Put`, and `Delete` methods to register routes.

```go
api := RestApi()

// Register the routes
api.Get("/users/:userId", ShowUser)
api.Put("/users/:userId", UpdateUser)
```

The `ShowUser` and `UpdateUser` functions are the handlers for the routes. You can define them as follows:

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

// A simple handler to update a user
func UpdateUser(request RestApiRequest) RestApiResponse {
	userId := request.PathParams["userId"]
	name := request.Body["name"]
	return RestApiResponse{
		Body: map[string]interface{}{
			"id":   userId,
			"name": name,
		},
		StatusCode: 200,
	}
}
```

## Local development

To run the application locally, you can use the following command:

```bash
$ go run ./main.go
```

## Testing

To run the tests, you can use the following command:

```bash
$ ./test.sh
```
