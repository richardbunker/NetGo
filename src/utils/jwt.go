package utils

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var JWTSecret []byte

type Session struct {
	GetClientId float64
	GetName string
}

func init() {
	// Load sample key data
	if keyData, e := os.ReadFile(".jwt_auth_key"); e == nil {
		JWTSecret = keyData
	} else {
		panic(e)
	}
}

// A function to generate a JWT
func GenerateJWT(clientId int, name string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})
	
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(JWTSecret)
	return tokenString, err
}

func ParseJWT(jwtString string) *Session {
	token, err := jwt.Parse(jwtString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return JWTSecret, nil
	})
	if err != nil {
		log.Println(err)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &Session{
			GetClientId: claims["clientId"].(float64),
			GetName: claims["name"].(string),
		}
	} else {
		fmt.Println(err)
		return nil
	}
}