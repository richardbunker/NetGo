package http

import (
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestRestApiHttpAdaptor(t *testing.T) {
	var r *http.Request
	var body io.Reader
	s := strings.NewReader(`{"name": "Mickey Mouse"}`)
	body = io.Reader(s)
	r, _ = http.NewRequest("GET", "/users/1234", body)
	r.Header.Add("Authorization", "Bearer")
	r.AddCookie(&http.Cookie{Name: "cookie1", Value: "value1"})
	adaptedRequest := StandardLibraryHTTPRequestAdaptor(r)

	if adaptedRequest.Method != "GET" {
		t.Errorf("Expected method to be 'GET', got %v", adaptedRequest.Method)
	}
	if adaptedRequest.Path != "/users/1234" {
		t.Errorf("Expected path to be '/users/1234', got %v", adaptedRequest.Path)
	}
	if adaptedRequest.Body["name"] != "Mickey Mouse" {
		t.Errorf("Expected body to be 'Mickey Mouse', got %v", adaptedRequest.Body["name"])
	}
	_, ok := adaptedRequest.Headers["Authorization"]
	if !ok {
		t.Errorf("Expected authorization header to be present, got %v", adaptedRequest.Headers)
	}

	if len(adaptedRequest.Cookies) != 1 {
		t.Errorf("Expected cookies to be empty, got %v", adaptedRequest.Cookies)
	}
	if adaptedRequest.Query != nil {
		t.Errorf("Expected query to be nil, got %v", adaptedRequest.Query)
	}
}
