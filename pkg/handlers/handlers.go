package handlers

import (
	"app/pkg/models"
	"encoding/json"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

func CreateUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var user models.User
	err := json.Unmarshal([]byte(request.Body), &user)

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       generateErrorBodyWithMessages("Invalid payload."),
		}, err
	}

	// TODO: save to database

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusCreated,
		Body:       user.ToJson(),
	}, err
}

func GetUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	user := models.User{
		Name: "Dummy GET User",
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       user.ToJson(),
	}, nil
}

func UpdateUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	user := models.User{
		Name: "Dummy Update User",
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       user.ToJson(),
	}, nil
}

func DeleteUser(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	user := models.User{
		Name: "Dummy Delete User",
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Body:       user.ToJson(),
	}, nil
}

func UnhandledMethod() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       generateErrorBodyWithMessages("Method not allowed."),
	}, nil
}

func generateErrorBodyWithMessages(message string) string {
	errorResponse := models.ErrorResponseBody{
		Message: message,
	}

	jbytes, err := json.Marshal(errorResponse)

	if err != nil {
		return "Something went wrong, please try again later."
	}

	return string(jbytes)
}
