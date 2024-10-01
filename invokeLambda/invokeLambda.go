package invokeLambda

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
	"github.com/energie-co-id/goLibrary/responseLib"
	"github.com/rs/zerolog/log"
)

type LambdaPayload struct {
	QueryStringParameters map[string]string `json:"queryStringParameters"`
	Body                  map[string]string `json:"body"`
	PathParameters        map[string]string `json:"pathParameters"`
}

func Invoke(lambdaName string, queryStringParameters map[string]string, body map[string]string, pathParameters map[string]string) (events.APIGatewayProxyResponse, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("ap-southeast-1"))
	if err != nil {
		log.Err(err).Msg("unable to load SDK config")
		responseLib.Generate(events.APIGatewayProxyRequest{}, 500, err.Error())
	}

	payload := LambdaPayload{
		QueryStringParameters: queryStringParameters,
		Body:                  body,
		PathParameters:        pathParameters,
	}

	payloadBytes, _ := json.Marshal(payload)

	svc := lambda.NewFromConfig(cfg)

	input := &lambda.InvokeInput{
		FunctionName: aws.String(lambdaName), // The name of the Lambda function to invoke
		Payload:      payloadBytes,           // Payload to pass to the invoked Lambda
	}

	result, err := svc.Invoke(context.TODO(), input)
	if err != nil {
		log.Err(err).Msg("failed to invoke lambda,")
		return responseLib.Generate(events.APIGatewayProxyRequest{}, int(result.StatusCode), err.Error())
	}
	response := events.APIGatewayProxyResponse{}
	err = json.Unmarshal(result.Payload, &response)
	if err != nil {
		log.Err(err).Msg("error marshal of response")
		return responseLib.Generate(events.APIGatewayProxyRequest{}, 500, err.Error())
	}
	return response, err
}
