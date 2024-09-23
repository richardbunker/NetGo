package middleware

import (
	"NetGo/services/auth"
	NetGoTypes "NetGo/types"
	"os"
	"testing"
)

func setupEnv(t *testing.T, key, value string) {
	err := os.Setenv(key, value)
	if err != nil {
		t.Fatalf("Failed to set environment variable: %v", err)
	}

	// Cleanup after the test to unset the environment variable
	t.Cleanup(func() {
		os.Unsetenv(key)
	})
}

func TestAuthorizePasses(t *testing.T) {
	// Mock environment variable
	setupEnv(t, "JWT_SECRET", "mocked_secret")

	validJwt, _ := auth.CreateJWT(NetGoTypes.User{
		Id:    "1",
		Email: "test@email.com",
		Name:  "Test User",
	}, []byte(os.Getenv("JWT_SECRET")), 1)
	authHeaders := []string{validJwt}
	request := NetGoTypes.RestApiRequest{
		Headers: map[string][]string{
			"Authorization": authHeaders,
		},
	}

	err, reason := Authenticated(request)

	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}
	if reason != nil {
		t.Errorf("Expected reason to be nil, got %v", reason)
	}
}
func TestAuthorizeRejectsWithoutAuthHeaderPresent(t *testing.T) {
	// Mock environment variable
	setupEnv(t, "JWT_SECRET", "mocked_secret")
	var authHeaders []string
	request := NetGoTypes.RestApiRequest{
		Headers: map[string][]string{
			"Authorization": authHeaders,
		},
	}

	err, reason := Authenticated(request)

	if err == nil {
		t.Errorf("Expected error to be 'Unauthenticated', got %v", err)
	}
	if reason == nil {
		t.Errorf("Expected reason to be nil, got %v", reason)
	}
	if reason.StatusCode != 401 {
		t.Errorf("Expected status code to be 401, got %v", reason.StatusCode)
	}
	if reason.Message != "Unauthenticated" {
		t.Errorf("Expected message to be 'Unauthenticated', got %v", reason.Message)
	}
}

func TestAuthorizeRejects(t *testing.T) {
	// Mock environment variable
	setupEnv(t, "JWT_SECRET", "mocked_secret")
	authHeaders := []string{"wrong_token"}
	request := NetGoTypes.RestApiRequest{
		Headers: map[string][]string{
			"Authorization": authHeaders,
		},
	}

	err, reason := Authenticated(request)

	if err == nil {
		t.Errorf("Expected error to be 'Unauthenticated', got %v", err)
	}
	if reason == nil {
		t.Errorf("Expected reason to be nil, got %v", reason)
	}
	if reason.StatusCode != 401 {
		t.Errorf("Expected status code to be 401, got %v", reason.StatusCode)
	}
	if reason.Message != "Unauthenticated" {
		t.Errorf("Expected message to be 'Unauthenticated', got %v", reason.Message)
	}
}
