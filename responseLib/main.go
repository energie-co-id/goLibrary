package responseLib

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type responseMessage struct {
	Message interface{} `json:"message"`
}

func Generate(event events.APIGatewayProxyRequest, statusCode int, message interface{}) (events.APIGatewayProxyResponse, error) {
	response, _ := json.Marshal(responseMessage{Message: message})
	return events.APIGatewayProxyResponse{
		Body:            string(response),
		StatusCode:      statusCode,
		IsBase64Encoded: false,
		Headers: map[string]string{
			"Content-Type":                     "application/json",
			"Access-Control-Allow-Origin":      "*",
			"Access-Control-Allow-Credentials": "true",
		},
	}, nil
}
