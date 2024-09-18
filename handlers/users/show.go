package users

import (
	. "NetGo/app"
	. "NetGo/db"
	. "NetGo/types"
)

// Show a user
func ShowUser(request RestApiRequest) RestApiResponse {
	// Query the database
	params := []interface{}{request.PathParams["userId"]}
	data := Query(`
		SELECT *
		FROM users
		WHERE id = ?
	`, params...)

	if len(data) == 0 {
		return ApiErrorResponse(404, "User not found")
	}

	// Create a new user
	user := User{
		Id:    data[0]["id"].(int64),
		Name:  data[0]["name"].(string),
		Email: data[0]["email"].(string),
	}
	// Return the user
	return RestApiResponse{
		StatusCode: 200,
		Body:       user,
	}
}
