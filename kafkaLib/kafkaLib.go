package kafkalib

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

func Producer(brokersUrl string, topic string, message string) error {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": brokersUrl})
	if err != nil {
		panic(err)
	}

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte(message),
	}, nil)

	// Wait for message deliveries before shutting down
	p.Flush(15 * 1000)

	defer p.Close()

	return err
}

func Consumer(brokersUrl string, topic string) (string, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": brokersUrl,
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{topic}, nil)

	// A signal handler or similar could be used to set this to false to break the loop.
	run := true
	message := ""
	for run {
		msg, err := c.ReadMessage(time.Second)
		if err == nil {
			message = string(msg.Value)
		} else if !err.(kafka.Error).IsTimeout() {
			// The client will automatically try to recover from all errors.
			// Timeout is not considered an error because it is raised by
			// ReadMessage in absence of messages.
			fmt.Printf("Consumer error: %v (%v)\n", err, msg)
		}
	}

	c.Close()

	return message, err
}
