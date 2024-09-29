package main

import (
	"NetGo/app"
	lib "NetGo/src/lib"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Load the environment variables from the .env file
	err := lib.LoadEnvFile(".env")
	if err != nil {
		log.Fatal("Error loading .env file: ", err)
	}
	api := app.Bootstrap()
	lib.StartUpMessage(os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), api); err != nil {
		fmt.Printf("Server failed to start: %s\n", err)
	}
}
