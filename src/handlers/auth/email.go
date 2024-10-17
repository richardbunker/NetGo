package auth

import (
	"NetGo/src/db/models"
	"NetGo/src/services/auth"
	"NetGo/src/services/notify"
	. "NetGo/src/types"
)

type EmailMagicLinkMessage struct {
	Status string `json:"status"`
}

// Email Magic Link
func EmailMagicLink(request NetGoRequest) NetGoResponse {
	userEmail, ok := request.Body["email"].(string)
	if !ok {
		return NetGoResponse{
			StatusCode: 400,
			Body:       "Invalid email",
		}
	}

	user, exists := models.FindUserByEmail(userEmail)
	if exists {
		token, err := auth.GenerateLoginToken(user.Id, 48, 1)
		if err == nil {
			mailer := notify.Mailer{}
			notify.EmailMagicLink(mailer, userEmail, token)
		}
	}
	return NetGoResponse{
		StatusCode: 200,
		Body:       EmailMagicLinkMessage{Status: "Done"},
	}
}
