package main

import (
	"app/pkg/handlers"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dynaClient *dynamodb.DynamoDB

const TABLE_NAME = "Users"

func main() {
	initializeDynamodbSession()
	lambda.Start(handler)
}

func initializeDynamodbSession() {
	log.Println("initialize dynamodb session")
	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials
	// and region from the shared configuration file ~/.aws/config.
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	log.Println("session created successfully")
	log.Println("connecting")

	// Create DynamoDB client
	dynaClient = dynamodb.New(sess)

	log.Println("session created successfully")
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	log.Println("New request with method: ", request.HTTPMethod)

	switch request.HTTPMethod {
	case "GET":
		return handlers.GetUser(request)
	case "POST":
		return handlers.CreateUser(request, TABLE_NAME, dynaClient)
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
