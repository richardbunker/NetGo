package auth

import (
	db "NetGo/src/db/models"
	NetGoTypes "NetGo/src/types"
	"fmt"
)

type RegisterResponse struct {
	Status string `json:"status"`
}

func Register(request NetGoTypes.NetGoRequest) NetGoTypes.NetGoResponse {
	userEmail, emailOk := request.Body["email"].(string)
	name, nameOk := request.Body["name"].(string)
	if !emailOk || !nameOk {
		return NetGoTypes.NetGoResponse{
			StatusCode: 400,
			Body:       "Invalid email or name",
		}
	}

	// Check if the user already exists
	_, exists := db.FindUserByEmail(userEmail)
	if exists {
		// Obfuscate this knowledge and return a 200
		return NetGoTypes.NetGoResponse{
			StatusCode: 200,
			Body:       RegisterResponse{Status: "Done"},
		}
	}

	// Create the user is the database
	_, err := db.RegisterUser(userEmail, name)
	if err != nil {
		fmt.Println(err)
	}
	return NetGoTypes.NetGoResponse{
		StatusCode: 200,
		Body:       RegisterResponse{Status: "Done"},
	}
}
