package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user User
	err := json.Unmarshal([]byte(request.Body), &user)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       generateErrorBodyWithMessages("Invalid payload."),
		}, err
	}

}

func generateErrorBodyWithMessages(message string) string {
	errorResponse := ErrorResponseBody{
		Message: message,
	}

	jbytes, err := json.Marshal(errorResponse)

	if err != nil {
		return "Something went wrong, please try again later."
	}

	return string(jbytes)
}

type User struct {
	ID              int      `json:"id"`
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Bio             string   `json:"bio"`
	PhoneNumber     string   `json:"phoneNumber"`
	SocialMediaURLs string   `json:"socialMediaUrls"`
	Templates       []string `json:"templates"`
}

type BusinessCard struct {
	QRCode string `json:"qrCode"`
	User   User   `json:"userInfo"`
}

type ErrorResponseBody struct {
	Message string `json:"message"`
}
