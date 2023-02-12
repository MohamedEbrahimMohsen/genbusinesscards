package user

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

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
		log.Printf("Got error while fetching user: %s\n", err)
		return nil, err
	}

	usr := new(User)
	err = dynamodbattribute.UnmarshalMap(result.Item, usr)
	if err != nil {
		log.Printf("Got error mapping fetched user: %s\n", err)
		return nil, err
	}

	return usr, nil
}

func FetchUsers(email string, tableName string, dynaClient *dynamodb.DynamoDB) (*[]User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.Scan(input)
	if err != nil {
		log.Printf("Got error while fetching all users: %s\n", err)
		return nil, err
	}

	users := new([]User)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, users)
	if err != nil {
		log.Printf("Got error mapping fetched user: %s\n", err)
		return nil, err
	}

	return users, nil
}

// https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/dynamo-example-create-table-item.html
func CreateUser(usr *User, tableName string, dynaClient *dynamodb.DynamoDB) (*User, error) {
	av, err := dynamodbattribute.MarshalMap(usr)
	if err != nil {
		log.Printf("Got error marshalling new user: %s\n", err)
		return nil, errors.New("got error marshalling new user")
	}

	currUser, err := FetchUser(usr.Email, tableName, dynaClient)
	if err != nil {
		return nil, err
	}

	if len(currUser.Email) != 0 {
		log.Printf("email already exist.")
		return nil, errors.New("email already exist")
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(input)
	if err != nil {
		log.Printf("Got error calling PutItem: %s\n", err)
		return nil, errors.New("got error calling PutItem")
	}

	log.Println("Successfully added '" + usr.Email + " to table " + tableName)
	return usr, nil
}

func UpdateUser(usr *User, tableName string, dynaClient *dynamodb.DynamoDB) (*User, error) {
	_, err := FetchUser(usr.Email, tableName, dynaClient)

	if err != nil {
		return nil, errors.New("no user registered with this email")
	}

	av, err := dynamodbattribute.MarshalMap(usr)
	if err != nil {
		log.Printf("Got error marshalling new user: %s\n", err)
		return nil, errors.New("got error while deserialzing the user's info")
	}

	dynaClient.PutItem(&dynamodb.PutItemInput{
		Item:      av,
		TableName: &tableName,
	})

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
		log.Printf("not found email for deleting: %v\n", email)
		return nil, errors.New("no user registered with this email")
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
		log.Printf("got error while deleting user: %s\n", err)
		return nil, err
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
}

type BusinessCard struct {
	QRCode string `json:"qrCode"`
	User   User   `json:"userInfo"`
}
