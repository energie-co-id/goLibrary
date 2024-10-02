package kafkaLib

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
)

func InitConfig(ClientID string, saslMechanism *string, saslUser *string, saslPassword *string, timeout int32) *sarama.Config {
	// Kafka configuration
	config := sarama.NewConfig()
	config.ClientID = ClientID // Optional
	if saslUser != nil && saslMechanism != nil && saslPassword != nil {
		config.Net.SASL.Enable = true // Enable SASL
		config.Net.SASL.Mechanism = sarama.SASLMechanism(*saslMechanism)
		config.Net.SASL.User = *saslUser
		config.Net.SASL.Password = *saslPassword
		config.Net.TLS.Enable = true
		config.Net.WriteTimeout = time.Duration(timeout) * time.Second
		// Enable TLS for secure connection if required
		// config.Net.TLS.Config = &tls.Config{
		// 	InsecureSkipVerify: true, // Use for testing; you can load proper certs for production
		// }
	}
	config.Producer.Return.Successes = true
	return config
}
func Producer(brokersUrl string, config *sarama.Config, topic string, key string, value string) (string, error) {
	// Create Kafka producer
	producer, err := sarama.NewSyncProducer([]string{brokersUrl}, config)
	if err != nil {
		log.Err(err).Msg("Error creating Kafka producer")
		return "", err
	}
	defer func() {
		if err := producer.Close(); err != nil {
			log.Err(err).Msg("Error closing Kafka producer")
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
		log.Err(err).Msg("Error producing message")
		return "", err
	}

	return fmt.Sprintf("Message sent to partition %d at offset %d", partition, offset), nil
}

func Consumer(brokersUrl string, ClientID string, topic string, partition int32, config *sarama.Config) {

	// Create Kafka consumer
	consumer, err := sarama.NewConsumer([]string{brokersUrl}, config)
	if err != nil {
		log.Err(err).Msg("Error creating Kafka consumer")
	}
	defer func() {
		if err := consumer.Close(); err != nil {
			log.Err(err).Msg("Error closing Kafka consumer")
		}
	}()

	// Trap SIGINT and SIGTERM to gracefully shutdown
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	// Consume messages
	partitionConsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		log.Err(err).Msg("Error creating partition consumer")
	}
	defer func() {
		if err := partitionConsumer.Close(); err != nil {
			log.Err(err).Msg("Error closing partition consumer")
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
