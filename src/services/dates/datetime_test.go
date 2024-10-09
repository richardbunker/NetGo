package dates

import (
	"testing"
	"time"
)

func TestCreateExpiresAtTimeFn(t *testing.T) {
	testingExpiry := CreateExpiresAtTime(0)
	now := time.Now()

	if testingExpiry.Unix() != now.Unix() {
		t.Errorf("Expected time to be identical, got %v and %v", testingExpiry, now)
	}
}

func TestHasExpired(t *testing.T) {
	expiredString := time.Now().Add(-time.Duration(1)).Format(time.RFC3339)
	result := HasExpired(expiredString)
	if !result {
		t.Errorf("Time should be expired")
	}
}

func TestConvertTimeToString(t *testing.T) {
	timeString := "2020-01-01T00:00:00Z"
	testingTime, _ := time.Parse(time.RFC3339, timeString)

	timeInStringFormat := convertTimeToString(testingTime)
	if timeInStringFormat != timeString {
		t.Errorf("Expected string to be: %s, got: %s", timeString, timeInStringFormat)
	}
}
