package auth

import (
	"NetGo/src/services/dates"
	NetGoTypes "NetGo/src/types"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// CreateJWT creates a new JWT token for the given user along with setting an expiration time.
func CreateJWT(user NetGoTypes.User, key []byte, exp int) (string, error) {
	expirationTime := dates.CreateExpiresAtTime(exp)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": user,
			"exp":  expirationTime.Unix(),
		})
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", fmt.Errorf("Invalid signing key")
	}
	return signedToken, nil
}

// VerifyJWT verifies a JWT token
func VerifyJWT(tokenString string, key []byte) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		log.Printf("Error parsing JWT token: %v", err)
		return nil, fmt.Errorf("Invalid token")
	}

	// Check if the token has expired
	exp := token.Claims.(jwt.MapClaims)["exp"]
	expTime := time.Unix(int64(exp.(float64)), 0)
	if time.Now().After(expTime) {
		log.Printf("Token has expired")
		return nil, fmt.Errorf("Token has expired")
	}

	// Token is valid
	return token, nil
}
