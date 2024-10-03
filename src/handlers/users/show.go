package users

import (
	db "NetGo/src/db/models"
	NetGoHttp "NetGo/src/services/http"
	NetGoTypes "NetGo/src/types"
)

// Show a user
func ShowUser(request NetGoTypes.RestApiRequest) NetGoTypes.RestApiResponse {
	userId := request.PathParams["userId"]

	// Find the user in the database by id
	item, ok := db.FindUserById(userId)
	if !ok {
		return NetGoHttp.ApiResponse(404, "User not found")
	}
	user := NetGoTypes.User{
		Id:    item.Id,
		Name:  item.Name,
		Email: item.Email,
	}
	// Return the user
	return NetGoTypes.RestApiResponse{
		StatusCode: 200,
		Body:       user,
	}
}
