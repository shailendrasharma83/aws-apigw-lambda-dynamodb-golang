package handlers

import (
	"aws-apigw-lambda-dynamodb-golang/pkg/shared"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"log"
	"net/http"
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Get(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	log.Print("Query String Params ", request.QueryStringParameters)

	return shared.ApiResponse(http.StatusOK, "Get Method ::  Lambda Version :: "+lambdacontext.FunctionVersion)
}

func Post(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)
	log.Print("Query String Params ", request.QueryStringParameters)

	return shared.ApiResponse(http.StatusOK, "Post Method ::  Lambda Version :: "+lambdacontext.FunctionVersion)
}

func UnhandledMethod() (events.APIGatewayProxyResponse, error) {
	return shared.ApiResponse(http.StatusMethodNotAllowed, "method Not allowed")
}
