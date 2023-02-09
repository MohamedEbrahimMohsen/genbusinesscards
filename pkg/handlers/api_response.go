package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func apiResponse(status int, body any) (*events.APIGatewayProxyResponse, error) {
	resp := events.APIGatewayProxyResponse{
		StatusCode: status,
		Headers:    map[string]string{"Content-Type": "application/json"},
	}

	strBody, _ := json.Marshal(body)
	resp.Body = string(strBody)
	return &resp, nil
}
