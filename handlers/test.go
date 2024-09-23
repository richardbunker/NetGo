package handlers

import (
	"NetGo/services/auth"
	. "NetGo/types"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

// A test handler for JWT validation
func Test(request RestApiRequest) RestApiResponse {
	requestToken := request.Headers["Authorization"]
	if requestToken[0] == "" {
		return RestApiResponse{
			StatusCode: 401,
			Body:       "Unauthorized",
		}
	}

	token, err := auth.VerifyJWT(requestToken[0], []byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return RestApiResponse{
			StatusCode: 401,
			Body:       "Invalid token",
		}
	}

	fmt.Println(token.Valid)
	fmt.Println(token.Claims.(jwt.MapClaims)["user"])
	exp := token.Claims.(jwt.MapClaims)["exp"]
	// Parse the expiration time
	expTime := time.Unix(int64(exp.(float64)), 0)
	// Check if the token has expired
	if time.Now().After(expTime) {
		fmt.Println("Token has expired")
	}
	fmt.Println("Token has not expired")
	return RestApiResponse{
		Body: map[string]interface{}{
			"message": "Test successful",
		},
		StatusCode: 200,
	}
}
