package http

import (
	NetGoTypes "NetGo/src/types"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-lambda-go/events"
)

func LambdaAPIGatewayHTTPRequestAdaptor(request events.APIGatewayV2HTTPRequest) NetGoTypes.RestApiRequest {
	// Adapt the headers from a map[string]string to a map[string][]string
	headers := request.Headers
	adaptedHeaders := make(map[string][]string)
	for key, value := range headers {
		adaptedHeaders[key] = []string{value}
	}
	// Unmarshal the request body
	var requestBody map[string]interface{}
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		fmt.Println("Error unmarshalling request body")
		requestBody = make(map[string]interface{})
	}
	adaptedRequest := NetGoTypes.RestApiRequest{
		Headers: adaptedHeaders,
		Cookies: request.Cookies,
		Method:  request.RequestContext.HTTP.Method,
		Path:    request.RequestContext.HTTP.Path,
		Query:   request.QueryStringParameters,
		Body:    requestBody,
	}

	return adaptedRequest
}
