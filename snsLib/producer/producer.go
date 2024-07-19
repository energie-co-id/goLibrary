package producer

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/energie-co-id/goLibrary/snsLib/models"
)

func PublishMessage(topicArn string, subject string, message models.SnsMessageType) (*sns.PublishOutput, error) {
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

	publishMessage, _ := json.Marshal(message)
	// Publish the message to the SNS topic
	result, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(string(publishMessage)),
		Subject:  aws.String(subject),
		TopicArn: aws.String(topicArn),
	})
	if err != nil {
		fmt.Println("Error publishing message:", err)
		return result, err
	}

	return result, nil
}
