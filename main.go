package main

import (
	"app/pkg/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	switch request.HTTPMethod {
	case "GET":
		return handlers.GetUser(request)
	case "POST":
		return handlers.CreateUser(request)
	case "PUT":
	case "PATCH":
		return handlers.UpdateUser(request)
	case "DELETE":
		return handlers.DeleteUser(request)
	default:
		return handlers.UnhandledMethod()
	}
	return handlers.UnhandledMethod()
}
