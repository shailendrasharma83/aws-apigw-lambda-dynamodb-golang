package main

import (
	"aws-apigw-lambda-dynamodb-golang/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch req.HTTPMethod {

	case "GET":
		return handlers.Get(req)
	case "POST":
		return handlers.Post(req)
	default:
		return handlers.UnhandledMethod()
	}
}
