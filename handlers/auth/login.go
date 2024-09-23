package auth

import (
	"NetGo/db/models"
	"NetGo/services/auth"
	"NetGo/services/dates"
	NetGoTypes "NetGo/types"
	"fmt"
	"os"
	"strconv"
)

type SuccessfulLoginResponse struct {
	JWT string `json:"jwt"`
}

// Login a user
func Login(request NetGoTypes.RestApiRequest) NetGoTypes.RestApiResponse {
	// Extract the token from the request query
	token := request.Body["token"].(string)

	// Find the login token in the database
	loginToken, err := models.FindLoginToken(token)
	if err != nil {
		return NetGoTypes.RestApiResponse{
			StatusCode: 401,
			Body:       "Invalid token",
		}
	}

	// Check if the token has expired
	if dates.HasExpired(loginToken.ExpiresAt) {
		// Delete the token from the database
		models.DeleteLoginToken(loginToken)
		return NetGoTypes.RestApiResponse{
			StatusCode: 401,
			Body:       "Token has expired",
		}
	}

	// Find the user in the database
	user, ok := models.FindUserById(loginToken.UserId)
	if !ok {
		return NetGoTypes.RestApiResponse{
			StatusCode: 404,
			Body:       "User not found",
		}
	}

	jwt := generateJWTForUser(user)

	// Delete the token from the database
	models.DeleteLoginToken(loginToken)

	return NetGoTypes.RestApiResponse{
		StatusCode: 200,
		Body:       SuccessfulLoginResponse{JWT: jwt},
	}
}

// A private helper function to generate a JWT token based off user data
func generateJWTForUser(user NetGoTypes.User) string {
	key := []byte(os.Getenv("JWT_SECRET"))
	expirationHourString := os.Getenv("JWT_EXPIRES_IN_N_HOURS")
	expirationTimeInt, _ := strconv.ParseInt(expirationHourString, 10, 64)
	token, err := auth.CreateJWT(user, key, int(expirationTimeInt))
	if err != nil {
		fmt.Println("Error generating JWT: ", err)
		return ""
	}
	return token
}
