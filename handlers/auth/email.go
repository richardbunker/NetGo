package auth

import (
	"NetGo/db/models"
	"NetGo/services/auth"
	"NetGo/services/notify"
	. "NetGo/types"
)

type EmailMagicLinkMessage struct {
	Status string `json:"status"`
}

// Email Magic Link
func EmailMagicLink(request RestApiRequest) RestApiResponse {
	userEmail := request.Body["email"].(string)
	user, exists := models.FindUserByEmail(userEmail)
	if exists {
		token, err := auth.GenerateLoginToken(user.Id, 48, 1)
		if err == nil {
			notify.EmailMagicLink(userEmail, token)
		}
	}
	return RestApiResponse{
		StatusCode: 200,
		Body:       EmailMagicLinkMessage{Status: "Done"},
	}
}
