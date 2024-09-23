package auth

import (
	db "NetGo/db/models"
	NetGoTypes "NetGo/types"
	"fmt"
)

type RegisterResponse struct {
	Status string `json:"status"`
}

func Register(request NetGoTypes.RestApiRequest) NetGoTypes.RestApiResponse {
	userEmail := request.Body["email"].(string)
	name := request.Body["name"].(string)

	// Check if the user already exists
	_, exists := db.FindUserByEmail(userEmail)
	if exists {
		// Obfuscate this knowledge and return a 200
		return NetGoTypes.RestApiResponse{
			StatusCode: 200,
			Body:       RegisterResponse{Status: "Done"},
		}
	}

	// Create the user is the database
	_, err := db.RegisterUser(userEmail, name)
	if err != nil {
		fmt.Println(err)
	}
	return NetGoTypes.RestApiResponse{
		StatusCode: 200,
		Body:       RegisterResponse{Status: "Done"},
	}
}
