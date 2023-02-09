package user

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func FetchUser(email string, tableName string, dynaClient *dynamodb.DynamoDB) (*User, error) {
	av, _ := dynamodbattribute.MarshalMap(email)
	result, err := dynaClient.GetItem(&dynamodb.GetItemInput{
		Key:       av,
		TableName: aws.String(tableName),
	})

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

	if currUser != nil {
		log.Printf("email already exist.")
		return nil, errors.New("email already exist.")
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

func UpdateUser() {

}

func DeleteUser() {
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
