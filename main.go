package main

import (
	. "NetGo/app"
	. "NetGo/handlers/users"
	. "NetGo/lib"
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

	// Create a new API application
	app := RestApi()
	app.Get("/users", RouteOptions{
		Handler: IndexUsers,
	})
	app.Get("/users/:userId", RouteOptions{
		Handler: ShowUser,
	})
	StartUpMessage(os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), app); err != nil {
		fmt.Printf("Server failed to start: %s\n", err)
	}
}
