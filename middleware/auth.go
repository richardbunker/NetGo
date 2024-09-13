package middleware

import (
	. "NetGo/types"
	"fmt"
	"os"
)

func getAuthToken() string {
	return "Bearer " + os.Getenv("AUTH_TOKEN")
}

func Authorize(request RestApiRequest) (error, *MiddlewareReason) {
	if request.Headers["authorization"][0] == getAuthToken() {
		return nil, nil
	}
	return fmt.Errorf("Unauthorized"), &MiddlewareReason{
		StatusCode: 401,
		Message:    "Unauthorized",
	}
}
