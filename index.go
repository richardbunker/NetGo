package main

import (
	"fmt"
	"net/http"
)

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/", helloWorldHandler)
	fmt.Println("✨ Server is listening on port 3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Printf("Server failed to start: %s\n", err)
	}
}
