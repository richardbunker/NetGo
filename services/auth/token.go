package auth

import (
	db "NetGo/db/models"
	"fmt"
	"math/rand"
	"time"
)

// This function gerenates a new login auth token and stores it in the database
func GenerateLoginToken(userId string, tokenLength int, expiresInNHours int) (string, error) {
	token := randomString(tokenLength)
	expiresAt := time.Now().Add(time.Hour * time.Duration(expiresInNHours)).Format(time.RFC3339)
	err := db.StoreLoginToken(userId, token, expiresAt)
	if err != nil {
		return "", fmt.Errorf("Error storing login token: %v", err)
	}
	return token, nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Create a new random generator
var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

// Generate a random string with a fixed length
func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
