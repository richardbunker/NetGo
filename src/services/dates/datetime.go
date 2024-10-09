package dates

import (
	"time"
)

// A helper function to convert a time to a RFC3339
func convertTimeToString(t time.Time) string {
	return t.Format(time.RFC3339)
}

func CreateExpiresAtTime(expiresInNHours int) time.Time {
	return time.Now().Add(time.Hour * time.Duration(expiresInNHours))
}

// HasExpired checks if the given time has expired.
func HasExpired(expiresAt string) bool {
	expirationTime, _ := time.Parse(time.RFC3339, expiresAt)
	return time.Now().After(expirationTime)
}
