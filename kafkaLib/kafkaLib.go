package kafkaLib

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/IBM/sarama"
)

func initConfig(ClientID string) *sarama.Config {
	// Kafka configuration
	config := sarama.NewConfig()
	config.ClientID = ClientID // Optional
	config.Producer.Return.Successes = true

	return config
}
func Producer(brokersUrl string, ClientID string, topic string, key string, value string) (string, error) {
	config := initConfig(ClientID)

	// Create Kafka producer
	producer, err := sarama.NewSyncProducer([]string{brokersUrl}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka producer: %v", err)
		return "", err
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Fatalf("Error closing Kafka producer: %v", err)
		}
	}()

	// Produce messages
	message := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(key),   // Change to your Kafka topic
		Value: sarama.StringEncoder(value), // Change to your message
	}
	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Printf("Error producing message: %v", err)
		return "", err
	}

	return fmt.Sprintf("Message sent to partition %d at offset %d", partition, offset), nil
}

func Consumer(brokersUrl string, ClientID string, topic string, partition int32) {
	config := initConfig(ClientID)

	// Create Kafka consumer
	consumer, err := sarama.NewConsumer([]string{brokersUrl}, config)
	if err != nil {
		log.Fatalf("Error creating Kafka consumer: %v", err)
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Fatalf("Error closing Kafka consumer: %v", err)
		}
	}()

	// Trap SIGINT and SIGTERM to gracefully shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Consume messages
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		log.Fatalf("Error creating partition consumer: %v", err)
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Fatalf("Error closing partition consumer: %v", err)
		}
	}()

	// Handle consumed messages
	go func() {
		for {
			select {
			case <-signals:
				log.Printf("<-signals")
				return
			case message := <-partitionConsumer.Messages():
				log.Printf("nunggu message")
				fmt.Printf("Received message: %s\n", string(message.Value))
			}
		}
	}()

	// Wait for termination signal
	<-signals
}
