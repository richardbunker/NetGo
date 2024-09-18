package users

import (
	. "NetGo/db"
	. "NetGo/types"
)

// Show a user
func IndexUsers(request RestApiRequest) RestApiResponse {
	// Query the database
	data := Query(`
		SELECT *
		FROM users
	`)

	results := []User{}

	for _, row := range data {
		user := User{
			Id:    row["id"].(int64),
			Name:  row["name"].(string),
			Email: row["email"].(string),
		}
		results = append(results, user)
	}

	// Return the user
	return RestApiResponse{
		StatusCode: 200,
		Body:       results,
	}
}
