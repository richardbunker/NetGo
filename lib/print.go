package lib

import (
	. "NetGo/types"
	"encoding/json"
	"fmt"
)

func StartUpMessage(portString string) {
	fmt.Print(`
    _   __     __  ______    
   / | / /__  / /_/ ____/___ 
  /  |/ / _ \/ __/ / __/ __ \
 / /|  /  __/ /_/ /_/ / /_/ /
/_/ |_/\___/\__/\____/\____/ 

`)
	fmt.Println("🚀 Launching NetGo...")
	fmt.Println()
	fmt.Printf("📡 Server is listening on port %s\n\n", portString)
}

func LogRequest(request RestApiRequest) {
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
