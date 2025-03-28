package users

import (
	db "NetGo/src/db/models"
	. "NetGo/src/types"
)

// Show a user
func IndexUsers(request NetGoRequest) NetGoResponse {
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
	return NetGoResponse{
		StatusCode: 200,
		Body:       users,
	}
}
