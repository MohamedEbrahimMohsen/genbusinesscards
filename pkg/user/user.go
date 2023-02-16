package user

import (
	"errors"
	"log"

	"app/pkg/codes"
	"app/pkg/services"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const BASE_URL = "https://gbc/profile/"

func FetchUser(email string, tableName string, dynaClient *dynamodb.DynamoDB) (*User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.GetItem(input)

	if err != nil {
		log.Printf("%s | %s\n", codes.E008, err)
		return nil, errors.New(codes.E008)
	}

	usr := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, usr)
	if err != nil {
		log.Printf("%s | %s\n", codes.E009, err)
		return nil, errors.New(codes.E009)
	}

	return usr, nil
}

func FetchUsers(email string, tableName string, dynaClient *dynamodb.DynamoDB) (*[]User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.Scan(input)
	if err != nil {
		log.Printf("%s | %s\n", codes.E010, err)
		return nil, errors.New(codes.E010)
	}

	users := new([]User)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, users)
	if err != nil {
		log.Printf("%s | %s\n", codes.E011, err)
		return nil, errors.New(codes.E011)
	}

	return users, nil
}

// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-create-table-item.html
func CreateUser(usr *User, tableName string, dynaClient *dynamodb.DynamoDB) (*User, error) {
	qr, err := services.GenerateQR(BASE_URL + usr.Email)
	if err != nil {
		return nil, err
	}

	usr.QRCode = qr
	item, err := dynamodbattribute.MarshalMap(usr)
	if err != nil {
		log.Printf("%s | %s\n", codes.E012, err)
		return nil, errors.New(codes.E012)
	}

	currUser, err := FetchUser(usr.Email, tableName, dynaClient)
	if err != nil {
		return nil, err
	}

	if len(currUser.Email) != 0 {
		log.Printf(codes.E013)
		return nil, errors.New(codes.E013)
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		log.Printf("%s | %s\n", codes.E014, err)
		return nil, errors.New(codes.E014)
	}

	return usr, nil
}

func UpdateUser(usr *User, tableName string, dynaClient *dynamodb.DynamoDB) (*User, error) {
	_, err := FetchUser(usr.Email, tableName, dynaClient)

	if err != nil {
		log.Printf(codes.E015)
		return nil, errors.New(codes.E015)
	}

	qr, err := services.GenerateQR(BASE_URL + usr.Email)
	if err != nil {
		return nil, err
	}

	usr.QRCode = qr
	av, err := dynamodbattribute.MarshalMap(usr)
	if err != nil {
		log.Printf("%s | %s\n", codes.E016, err)
		return nil, errors.New(codes.E016)
	}

	_, err = dynaClient.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: &tableName,
	})

	if err != nil {
		log.Printf("%s | %s\n", codes.E017, err)
		return nil, errors.New(codes.E017)
	}

	return usr, nil
}

// https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/LegacyConditionalParameters.AttributeUpdates.html
// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-update-table-item.html
func PatchUpdateUser(usr *User, tableName string, dynaClient *dynamodb.DynamoDB) (*User, error) {
	// NOT IMPLEMENTED!
	return UpdateUser(usr, tableName, dynaClient)
}

// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-delete-table-item.html
func DeleteUser(email string, tableName string, dynaClient *dynamodb.DynamoDB) (*User, error) {
	usr, err := FetchUser(email, tableName, dynaClient)

	if err != nil {
		log.Printf(codes.E018)
		return nil, errors.New(codes.E018)
	}

	_, err = dynaClient.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"email": {
				S: aws.String(email),
			},
		},
		TableName: &tableName,
	})

	if err != nil {
		log.Printf("%s | %s\n", codes.E019, err)
		return nil, errors.New(codes.E019)
	}

	return usr, nil
}

type User struct {
	Name            string   `json:"name"`
	Email           string   `json:"email"`
	Bio             string   `json:"bio"`
	PhoneNumber     string   `json:"phoneNumber"`
	SocialMediaURLs string   `json:"socialMediaUrls"`
	Templates       []string `json:"templates"`
	QRCode          []byte   `json:"qrCode"`
}
