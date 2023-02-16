package handlers

import (
	"app/pkg/codes"
	"app/pkg/user"
	"app/pkg/validators"
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
		log.Printf("%s | %s\n", codes.E001, err)
		return apiResponse(http.StatusBadRequest, ErrorBody{ErrorMsg: aws.String(codes.E001)})
	}

	if !validators.IsEmailValid(usr.Email) {
		log.Printf(codes.E002)
		return apiResponse(http.StatusBadRequest, ErrorBody{ErrorMsg: aws.String(codes.E002)})
	}

	_, err = user.CreateUser(&usr, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{ErrorMsg: aws.String(err.Error())})
	}

	log.Println("User created successfully!")
	return apiResponse(http.StatusCreated, usr)
}

func GetUser(request events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB) (*events.APIGatewayProxyResponse, error) {
	log.Println("Fetching user...")
	email := request.QueryStringParameters["email"]

	if len(email) > 0 {
		usr, err := user.FetchUser(email, tableName, dynaClient)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{ErrorMsg: aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, &usr)
	}

	usr, err := user.FetchUsers(email, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{ErrorMsg: aws.String(err.Error())})
	}

	log.Println("User fetched successfully!")
	return apiResponse(http.StatusOK, &usr)
}

func UpdateUser(request events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB) (*events.APIGatewayProxyResponse, error) {
	log.Println("Updating user...")
	var usr = new(user.User)

	err := json.Unmarshal([]byte(request.Body), usr)
	if err != nil {
		log.Printf("%s | %s\n", codes.E003, err)
		return apiResponse(http.StatusBadRequest, ErrorBody{ErrorMsg: aws.String(codes.E003)})
	}

	usr, err = user.UpdateUser(usr, tableName, dynaClient)

	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	log.Println("User Updated successfully!")
	return apiResponse(http.StatusOK, usr)
}

func PatchUpdateUser(request events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB) (*events.APIGatewayProxyResponse, error) {
	log.Println("Patch updating user...")
	var usr = new(user.User)

	err := json.Unmarshal([]byte(request.Body), usr)
	if err != nil {
		log.Printf("%s | %s\n", codes.E004, err)
		return apiResponse(http.StatusBadRequest, ErrorBody{ErrorMsg: aws.String(codes.E004)})
	}

	usr, err = user.PatchUpdateUser(usr, tableName, dynaClient)

	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	log.Println("User patched successfully!")
	return apiResponse(http.StatusOK, usr)
}

func DeleteUser(request events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.DynamoDB) (*events.APIGatewayProxyResponse, error) {
	log.Println("Deleting user...")
	email := request.QueryStringParameters["email"]

	if len(email) == 0 {
		log.Println(codes.E005)
		return apiResponse(http.StatusBadRequest, ErrorBody{ErrorMsg: aws.String(codes.E005)})
	}

	usr, err := user.DeleteUser(email, tableName, dynaClient)

	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}

	log.Println("User deleted successfully!")
	return apiResponse(http.StatusOK, usr)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	log.Println(codes.E006)
	return apiResponse(http.StatusMethodNotAllowed, ErrorBody{ErrorMsg: aws.String(codes.E006)})
}
