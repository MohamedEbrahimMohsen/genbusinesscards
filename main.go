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
		return handlers.GetUser(request, TABLE_NAME, dynaClient)
	case "POST":
		return handlers.CreateUser(request, TABLE_NAME, dynaClient)
	case "PUT":
		return handlers.UpdateUser(request, TABLE_NAME, dynaClient)
	case "PATCH":
		return handlers.PatchUpdateUser(request, TABLE_NAME, dynaClient)
	case "DELETE":
		return handlers.DeleteUser(request, TABLE_NAME, dynaClient)
	default:
		return handlers.UnhandledMethod()
	}
}
