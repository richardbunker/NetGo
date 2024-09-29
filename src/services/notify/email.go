package notify

import (
	"fmt"
	"net/smtp"
	"os"
)

func EmailMagicLink(email string, token string) {
	message := fmt.Sprintf("Please click here to login: https://example.com/login?token=%s", token)
	SendEmail(email, "Magic Link", []byte(message))
}

// Send an email
func SendEmail(to string, subject string, message []byte) {
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
	smtpHost := os.Getenv("EMAIL_SMTP_HOST")
	smtpPort := os.Getenv("EMAIL_SMTP_PORT")

	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Create email message with headers
	body := fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s\r\n",
		from, to, subject, string(message),
	)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{to}, []byte(body))
	if err != nil {
		fmt.Println("Error sending email: ", err)
	}

	fmt.Printf("\n📮 Email has been successfully sent to: %s\n", to)
}
