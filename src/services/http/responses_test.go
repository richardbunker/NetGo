package http

import (
	"encoding/json"
	"testing"
)

func TestApiResponseReturnsErrorFormat(t *testing.T) {
	testingResponse := ApiResponse(401, "Unauthenticated")
	if testingResponse.Body != nil {
		stringed, _ := json.Marshal(testingResponse.Body)
		if string(stringed) != "{\"error\":\"Unauthenticated\"}" {
			t.Errorf("Expected, got %s", stringed)
		}
	}
}

func TestApiResponseReturnsStandardFormat(t *testing.T) {
	testingResponse := ApiResponse(201, "Created")
	if testingResponse.Body != nil {
		stringed, _ := json.Marshal(testingResponse.Body)
		if string(stringed) != "{\"message\":\"Created\"}" {
			t.Errorf("Expected, got %s", stringed)
		}
	}
}
