package http

import (
	NetGoTypes "NetGo/src/types"
)

func ApiResponse(statusCode int, message string) NetGoTypes.RestApiResponse {
	var key string
	if statusCode >= 400 {
		key = "error"
	} else {
		key = "message"
	}
	return NetGoTypes.RestApiResponse{
		Body:       map[string]string{key: message},
		StatusCode: statusCode,
	}
}
