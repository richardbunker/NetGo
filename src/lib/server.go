package lib

import (
	"NetGo/src/console"
	"NetGo/src/routes"
	"fmt"
	"net/http"
)

func NetGoServer(port int) {
	portString := fmt.Sprintf("%d", port)
	console.PrettyBoot(portString)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Router(w, r, routes.ApiRoutes)
	})
	if err := http.ListenAndServe(":" + portString, nil); err != nil {
		fmt.Printf("Server failed to start: %s\n", err)
	}
}