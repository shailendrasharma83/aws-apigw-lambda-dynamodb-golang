package main

import (
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	log.Print("Query String Params ", request.QueryStringParameters)

	return events.APIGatewayProxyResponse{
		Body:       "Hello From Lambda Version :: " + lambdacontext.FunctionVersion,
		StatusCode: 200,
	}, nil
}

func main() {
	lambda.Start(Handler)
}
