package users

import (
	db "NetGo/db/models"
	. "NetGo/types"
)

// Show a user
func IndexUsers(request RestApiRequest) RestApiResponse {
	// Get all users
	items, _ := db.ListUsers()

	users := []User{}
	for _, item := range items {
		users = append(users, User{
			Id:    item.Id,
			Name:  item.Name,
			Email: item.Email,
		})
	}
	// Return the user
	return RestApiResponse{
		StatusCode: 200,
		Body:       users,
	}
}
