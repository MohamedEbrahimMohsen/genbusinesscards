package user

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func FetchUser() {

}

func FetchUsers() {

}

func CreateUser(usr *User, tableName string, dynaClient *dynamodb.DynamoDB) (*User, error) {

	av, err := dynamodbattribute.MarshalMap(usr)
	if err != nil {
		log.Printf("Got error marshalling new user: %s\n", err)
		return nil, errors.New("got error marshalling new user")
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
