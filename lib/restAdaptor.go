package lib

import (
	"encoding/json"
	"net/http"
	. "rest_api/types"
)

func RestApiAdaptor(r *http.Request) RestApiRequest {
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
