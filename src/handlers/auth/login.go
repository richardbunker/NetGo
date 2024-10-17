package auth

import (
	"NetGo/src/db/models"
	"NetGo/src/services/auth"
	"NetGo/src/services/dates"
	NetGoTypes "NetGo/src/types"
	"fmt"
	"os"
	"strconv"
)

type SuccessfulLoginResponse struct {
	JWT string `json:"jwt"`
}

// Login a user
func Login(request NetGoTypes.NetGoRequest) NetGoTypes.NetGoResponse {
	// Extract the token from the request query
	token, ok := request.Body["token"].(string)
	if !ok {
		return NetGoTypes.NetGoResponse{
			StatusCode: 400,
			Body:       "Invalid token",
		}
	}

	// Find the login token in the database
	loginToken, err := models.FindLoginToken(token)
	if err != nil {
		return NetGoTypes.NetGoResponse{
			StatusCode: 401,
			Body:       "Invalid token",
		}
	}

	// Check if the token has expired
	if dates.HasExpired(loginToken.ExpiresAt) {
		// Delete the token from the database
		models.DeleteLoginToken(loginToken)
		return NetGoTypes.NetGoResponse{
			StatusCode: 401,
			Body:       "Token has expired",
		}
	}

	// Find the user in the database
	user, ok := models.FindUserById(loginToken.UserId)
	if !ok {
		return NetGoTypes.NetGoResponse{
			StatusCode: 404,
			Body:       "User not found",
		}
	}

	jwt := generateJWTForUser(user)

	// Delete the token from the database
	models.DeleteLoginToken(loginToken)

	return NetGoTypes.NetGoResponse{
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
