package app

import (
	. "NetGo/src/api"
	. "NetGo/src/db"
	. "NetGo/src/handlers/auth"
	. "NetGo/src/handlers/users"
	"NetGo/src/middleware"
	. "NetGo/src/types"
	"log"
)

func Bootstrap() *Api {
	// Initialize the database singleton
	dbErr := NewDynamoDBSingleton()
	if dbErr != nil {
		log.Fatal("Error initializing DynamoDB: ", dbErr)
	}
	// Create a new API application
	api := NetGo()
	api.UseMiddleware([]Middleware{middleware.LogRequests})

	// Register the API Routes

	// ****************************************************
	// *************** 🔐 AUTHENTICATION ******************
	// ****************************************************
	/*
		This route sends a magic link to the user's email address.
		@method POST
		@route /auth/email-magic-link
	*/
	api.Post("/auth/email-magic-link", RouteOptions{
		Handler: EmailMagicLink,
	})
	/*
		This route logs the user in.
		@method POST
		@route /auth/login
	*/
	api.Post("/auth/login", RouteOptions{
		Handler: Login,
	})

	/*
		This route registers a new user.
		@method POST
		@route /auth/register
	*/
	api.Post("/auth/register", RouteOptions{
		Handler: Register,
	})

	// ****************************************************
	// *************** 👱 USERS ***************************
	// ****************************************************
	api.Group("/users", []Middleware{
		middleware.Authenticated,
	}, func() {
		/*
			List all users.
			@method GET
			@route /users
		*/
		api.Get("/", RouteOptions{
			Handler: IndexUsers,
		})
		/*
			Show a user.
			@method GET
			@route /users/:userId
		*/
		api.Get("/:userId", RouteOptions{
			Handler: ShowUser,
		})
	})

	// Return the api
	return api
}
