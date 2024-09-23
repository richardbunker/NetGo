package middleware

import (
	"NetGo/services/auth"
	. "NetGo/types"
	"fmt"
	"os"
)

func Authenticated(request RestApiRequest) (error, *MiddlewareReason) {
	requestToken := request.Headers["Authorization"]
	// If no authorization header is present, then return unauthorized
	if len(requestToken) == 0 {
		return unauthenticatedResponse()
	}

	// Verify the JWT token
	_, err := auth.VerifyJWT(requestToken[0], []byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return unauthenticatedResponse()
	}

	return nil, nil
}

// Unauthenticated middleware response
func unauthenticatedResponse() (error, *MiddlewareReason) {
	return fmt.Errorf("Unauthenticated"), &MiddlewareReason{
		StatusCode: 401,
		Message:    "Unauthenticated",
	}
}
