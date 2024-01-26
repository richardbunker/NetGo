package server

import (
	"NetGo/src/console"
	"NetGo/src/routes"
	"fmt"
	"net/http"
)

func Start(port int) {
	console.PrettyBoot()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		routes.Router(w, r)
	})
	portString := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(portString, nil); err != nil {
		fmt.Printf("Server failed to start: %s\n", err)
	}
}