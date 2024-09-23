package dates

import (
	"time"
)

// CreateExpiresAtString creates a string representation of the time that is n hours from now.
func CreateExpiresAtString(expiresInNHours int) string {
	return time.Now().Add(time.Hour * time.Duration(expiresInNHours)).Format(time.RFC3339)
}

func CreateExpiresAtTime(expiresInNHours int) time.Time {
	return time.Now().Add(time.Hour * time.Duration(expiresInNHours))
}

// HasExpired checks if the given time has expired.
func HasExpired(expiresAt string) bool {
	expirationTime, _ := time.Parse(time.RFC3339, expiresAt)
	return time.Now().After(expirationTime)
}
