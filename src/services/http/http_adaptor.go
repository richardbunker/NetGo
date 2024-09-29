package http

import (
	. "NetGo/src/types"
	"encoding/json"
	"net/http"
)

func StandardLibraryHTTPRequestAdaptor(r *http.Request) RestApiRequest {
	var cookies []string
	for _, cookie := range r.Cookies() {
		cookies = append(cookies, cookie.String())
	}
	body := make([]byte, r.ContentLength)
	r.Body.Read(body)
	bodyToMap := make(map[string]interface{})
	json.Unmarshal(body, &bodyToMap)

	adaptedRequest := RestApiRequest{
		Headers: r.Header,
		Cookies: cookies,
		Method:  r.Method,
		Path:    r.URL.Path,
		Body:    bodyToMap,
	}
	return adaptedRequest
}
