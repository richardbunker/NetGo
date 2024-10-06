package notify

import (
	"os"
	"testing"
)

// Mock environment variables for email sending
func init() {
	os.Setenv("APP_LOGIN_URL", "http://example.com/login")
}

type MockEmailSender struct {
	SentEmails []SentEmail
}

type SentEmail struct {
	To      string
	Subject string
	Message []byte
}

// Mock implementation of EmailSender
func (m *MockEmailSender) SendEmail(to string, subject string, message []byte) {
	m.SentEmails = append(m.SentEmails, SentEmail{To: to, Subject: subject, Message: message})
}

// TestEmailMagicLink tests the EmailMagicLink function
func TestEmailMagicLink(t *testing.T) {
	mockSender := &MockEmailSender{}
	EmailMagicLink(mockSender, "test@example.com", "dummy-token")

	if len(mockSender.SentEmails) != 1 {
		t.Errorf("expected 1 email to be sent, got %d", len(mockSender.SentEmails))
	}

	sentEmail := mockSender.SentEmails[0]
	if sentEmail.To != "test@example.com" {
		t.Errorf("expected recipient to be test@example.com, got %s", sentEmail.To)
	}

	if sentEmail.Subject != "Magic Link" {
		t.Errorf("expected subject to be 'Magic Link', got %s", sentEmail.Subject)
	}

	expectedMessage := "Please click here to login: http://example.com/login?token=dummy-token"
	if string(sentEmail.Message) != expectedMessage {
		t.Errorf("expected message to be '%s', got '%s'", expectedMessage, string(sentEmail.Message))
	}
}
