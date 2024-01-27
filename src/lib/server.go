package lib

import (
	"NetGo/src/console"
	"NetGo/src/routes"
	"NetGo/src/utils"
	"fmt"
	"net/http"
)

func NetGoServer(port int) {
	portString := fmt.Sprintf("%d", port)
	console.PrettyBoot(portString)
	http.HandleFunc("/auth/signin", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println("Requesting a JWT...")
			tokenString, err := utils.GenerateJWT(123456, "John Doe")	
			if err != nil {
				fmt.Println("Error generating JWT. Error:", err)
			}
			fmt.Println(tokenString)
		}
	})
	http.HandleFunc("/auth/validate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println("Validating a JWT...")
			claims := utils.ParseJWT("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjbGllbnRJZCI6MTIzNDU2LCJleHAiOjE3MDY0NDIwMjgsImlhdCI6MTcwNjM1NTYyOCwiaXNzIjoiTmV0R28iLCJuYW1lIjoiSm9obiBEb2UifQ.cR1o-zG4xBrSiHRxs7ZXbRTjanlQfpsS3js6l584RXE")
			fmt.Println(claims)
		}
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Router(w, r, routes.ApiRoutes)
	})
	if err := http.ListenAndServe(":" + portString, nil); err != nil {
		fmt.Printf("Server failed to start: %s\n", err)
	}
}