package auth

import (
	NetGoTypes "NetGo/types"
	"os"
	"testing"
)

var invalidMockedJWT string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3Mjc1MTg0MzcsInVzZXIiOnsiVXNlcklkIjoiNjU0MzIxIiwiRW1haWwiOiJ0ZXN0QGVtYWlsLmNvbSIsIk5hbWUiOiJKb2huIFNtaXRoIn19.zhjyTe0mt29kWZSqsroYn6xx9i6QJhfl1ojF5Z_ocrU"

var expiredMockedJWT string = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjczODc3MjEsInVzZXIiOnsiVXNlcklkIjoiMSIsIkVtYWlsIjoidGVzdEBlbWFpbC5jb20iLCJOYW1lIjoiVGVzdCBVc2VyIn19.zxGMeg2EYexYmhqTrpUqvn-lpu4r4y3494E31nzIjWI"

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

func TestGenerateJWT(t *testing.T) {
	// Mock environment variable
	setupEnv(t, "JWT_SECRET", "mocked_secret")

	validJwt, err := CreateJWT(NetGoTypes.User{
		Id:    "1",
		Email: "test@email.com",
		Name:  "Test User",
	}, []byte(os.Getenv("JWT_SECRET")), 1)

	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}

	if validJwt == "" {
		t.Errorf("Expected non-empty string, %v", validJwt)
	}
}

func TestVerifyJWT(t *testing.T) {
	// Mock environment variable
	setupEnv(t, "JWT_SECRET", "mocked_secret")
	validJwt, _ := CreateJWT(NetGoTypes.User{
		Id:    "1",
		Email: "test@email.com",
		Name:  "Test User",
	}, []byte(os.Getenv("JWT_SECRET")), 1)

	token, err := VerifyJWT(validJwt, []byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		t.Errorf("Expected error to be nil, got %v", err)
	}

	if token == nil {
		t.Errorf("Expected non-empty string, %v", token)
	}
}

func TestVerifyJWTHandlesMalformedTokens(t *testing.T) {
	// Mock environment variable
	setupEnv(t, "JWT_SECRET", "mocked_secret")

	token, err := VerifyJWT("some-malformed-token", []byte(os.Getenv("JWT_SECRET")))

	if err == nil {
		t.Errorf("Expected error 'Invalid token', got %v", err)
	}

	if token != nil {
		t.Errorf("Expected token to be nil, got %v", token)
	}
}

func TestVerifyJWTHandlesExpiredTokens(t *testing.T) {
	// Mock environment variable
	setupEnv(t, "JWT_SECRET", "mocked_secret")

	token, err := VerifyJWT(expiredMockedJWT, []byte(os.Getenv("JWT_SECRET")))

	if err == nil {
		t.Errorf("Expected error 'Invalid token', got %v", err)
	}

	if token != nil {
		t.Errorf("Expected token to be nil, got %v", token)
	}
}
