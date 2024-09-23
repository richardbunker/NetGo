package main

import (
	. "NetGo/api"
	. "NetGo/db"
	. "NetGo/handlers/auth"
	. "NetGo/handlers/users"
	. "NetGo/lib"
	"NetGo/middleware"
	. "NetGo/types"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load the environment variables
	err := LoadEnvFile(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}

	// Initialize the database singleton
	dbErr := NewDynamoDBSingleton()
	if dbErr != nil {
		log.Fatal("Error initializing DynamoDB: ", dbErr)
	}
	// Create a new API application
	api := RestApi()
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

	StartUpMessage(os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), api); err != nil {
		fmt.Printf("Server failed to start: %s\n", err)
	}
}
