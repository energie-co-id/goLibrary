package mqttLib

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"
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

func isMap(publishMessage interface{}) bool {
	_, ok1 := publishMessage.(map[string]interface{})
	_, ok2 := publishMessage.(map[string]string)
	return ok1 || ok2
}

func publish(client mqtt.Client, publishMessage interface{}, publishTopic string) {
	token := client.Publish(publishTopic, 0, false, publishMessage)
	token.Wait()
	// time.Sleep(time.Second)
}
func subscribe(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic: %s", topic)
}

func InitializeClient(brokerURL string, port int, clientID string, username string, password string) mqtt.Client {
	//config connect to mqtt
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("%s:%d", brokerURL, port))
	opts.SetClientID(clientID)
	opts.SetUsername(username)
	opts.SetPassword(password)
	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	return client
}

func MainPublish(client mqtt.Client, publishTopic string, publishMessage interface{}) (interface{}, error) {
	//check message
	if isMap(publishMessage) || reflect.TypeOf(publishMessage).Kind() == reflect.Struct { //if publish message type is object, we parse it to string
		// Convert the message to JSON before publishing
		publishMessageStr, err := json.Marshal(publishMessage)
		if err != nil {
			return nil, err
		}
		publishMessage = string(publishMessageStr)
	}
	//publish message
	publish(client, publishMessage, publishTopic)

	fmt.Printf("mqtt published, topic : %s, message: %v\n", publishTopic, publishMessage)
	client.Disconnect(250)

	return publishMessage, nil
}

func PublishAndListening(client mqtt.Client, publishTopic string, listeningTopic string, publishMessage interface{}) string {
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
