package producer

import (
	// "encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

// type snsMessageType struct {
// 	From          string `json:"from"`
// 	To            string `json:"to"`
// 	TemplateValue string `json:"templateValue"`
// }

func PublishMessage(topicArn string, subject string, message string) {
	// Create a session using shared config
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-southeast-1"),
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		return
	}

	// Create SNS client
	svc := sns.New(sess)

	// Define the message and the topic ARN
	// templateValueString, _ := json.Marshal(templateValue)
	// message2 := snsMessageType{
	// 	From:          "evgate-support@stroomer.id",
	// 	To:            "kevin@stroomer.id",
	// 	TemplateValue: string(templateValueString),
	// }
	// message2_, _ := json.Marshal(message2)
	// Publish the message to the SNS topic

	result, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(message),
		Subject:  aws.String(subject),
		TopicArn: aws.String(topicArn),
	})
	if err != nil {
		fmt.Println("Error publishing message:", err)
		return
	}

	// Print the message ID
	fmt.Println("Message ID:", *result.MessageId)
}
