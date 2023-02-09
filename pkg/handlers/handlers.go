package handlers

import (
	"app/pkg/user"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func CreateUser(request events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB) (*events.APIGatewayProxyResponse, error) {
	log.Println("Creating new user...")
	var usr user.User

	err := json.Unmarshal([]byte(request.Body), &usr)
	if err != nil {
		log.Printf("Got error while unmarshalling the request body: %s\n", err)
		return apiResponse(http.StatusBadRequest, ErrorBody{ErrorMsg: aws.String(err.Error())})
	}

	_, err = user.CreateUser(&usr, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{ErrorMsg: aws.String(err.Error())})
	}

	return apiResponse(http.StatusCreated, usr)
}

func GetUser(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	user := user.User{
		Name: "Dummy GET User",
	}

	return apiResponse(http.StatusOK, user)
}

func UpdateUser(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	user := user.User{
		Name: "Dummy Update User",
	}

	return apiResponse(http.StatusOK, user)
}

func DeleteUser(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	user := user.User{
		Name: "Dummy Delete User",
	}

	return apiResponse(http.StatusOK, user)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorBody{ErrorMsg: aws.String("Method Not Allowed")})
}
