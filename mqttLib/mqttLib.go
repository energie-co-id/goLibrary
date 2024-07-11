package mqttLib

import (
	"fmt"
	"log"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var receivedMessage string = "no response"

var messagePubHandler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	receivedMessage = string(msg.Payload())
	fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler mqtt.OnConnectHandler = func(client mqtt.Client) {
	log.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func publish(client mqtt.Client, publishMessage string, publishTopic string) {
	token := client.Publish(publishTopic, 0, false, publishMessage)
	token.Wait()
	// time.Sleep(time.Second)
}
func subscribe(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}

func mqttConfig() mqtt.Client {
	//config connect to mqtt
	var broker = os.Getenv("MQTT_URL")
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(os.Getenv("MQTT_USERNAME"))
	opts.SetPassword(os.Getenv("MQTT_PASS"))
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}

func MainPublish(publishMessage string, publishTopic string) (string, error) {
	//config connect to mqtt
	client := mqttConfig()

	//publish message
	publish(client, publishMessage, publishTopic)

	client.Disconnect(250)

	return publishMessage, nil
}

func PublishAndListening(publishMessage string, publishTopic string, listeningTopic string) string {
	client := mqttConfig()

	subscribe(client, listeningTopic)
	publish(client, publishMessage, publishTopic)
	targetTime := 10 * time.Second // 10 seconds
	timer := 0
	for receivedMessage == "no response" && timer < int(targetTime) {
		time.Sleep(time.Millisecond) // Wait for 1 millisecond
		timer += int(time.Millisecond)
	}

	client.Disconnect(250)

	response := receivedMessage
	receivedMessage = "no response"
	return response
}
