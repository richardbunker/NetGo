package main

import (
	"NetGo/app"
	"NetGo/src/services/http"
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	req := http.LambdaAPIGatewayHTTPRequestAdaptor(request)
	api := app.Bootstrap()
	res := api.HandleRequest(req)
	body, _ := json.Marshal(res.Body)
	return events.APIGatewayV2HTTPResponse{
		StatusCode: 200,
		Body:       string(body),
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
