package middleware

import (
	. "NetGo/types"
	"encoding/json"
	"fmt"
)

func LogRequests(request RestApiRequest) (error, *MiddlewareReason) {
	logRequest(request)
	return nil, nil
}

// Log the incoming request
func logRequest(request RestApiRequest) {
	infoLine := ""
	headersTitle := "🔐 Headers: "
	headers := []string{}
	for key, value := range request.Headers {
		headers = append(headers, fmt.Sprintf("%s -> %v", key, value))
	}
	cookiesTitle := "🍪 Cookies: "
	cookies := []string{}
	for _, cookie := range request.Cookies {
		for _, value := range cookie {
			cookies = append(cookies, fmt.Sprintf("%v", value))
		}
	}
	bodyTitle := "📦 Body: "
	body := []string{}
	for key, value := range request.Body {
		body = append(body, fmt.Sprintf("%s -> %v", key, value))
	}
	if len(headers) > 0 {
		headerString := ""
		for _, header := range headers {
			headerString += header + ", "
		}
		infoLine += headersTitle + fmt.Sprintf("%s", headerString)
	}
	if len(cookies) > 0 {
		infoLine += " " + cookiesTitle + fmt.Sprintf("%v", cookies)
	}
	if len(body) > 0 {
		jsonBody, _ := json.Marshal(request.Body)
		infoLine += " " + bodyTitle + fmt.Sprintf("%s", jsonBody)
	}

	fmt.Printf("\n📬 Incoming Request: '[%s]%s' %s \n", request.Method, request.Path, infoLine)
}
