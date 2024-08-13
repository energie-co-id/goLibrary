package producer

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

func PublishMessage(topicArn string, subject string, message string) (*sns.PublishOutput, error) {
	// Create a session using shared config
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return nil, err
	}

	// Create SNS client
	svc := sns.New(sess)
	// Publish the message to the SNS topic
	result, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(message),
		Subject:  aws.String(subject),
		TopicArn: aws.String(topicArn),
	})
	if err != nil {
		fmt.Println("Error publishing message:", err)
		return result, err
	}

	return result, nil
}
