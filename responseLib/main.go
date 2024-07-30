package responseLib

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type responseMessage struct {
	Message interface{} `json:"message"`
}

func Generate(event events.APIGatewayProxyRequest, statusCode int, message interface{}) (events.APIGatewayProxyResponse, error) {
	// Check if the message is of type error and convert it to a string if so
	if err, ok := message.(error); ok {
		message = err.Error()
	}

	response, err := json.Marshal(responseMessage{Message: message})
	if err != nil {
		return events.APIGatewayProxyResponse{}, err
	}

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
