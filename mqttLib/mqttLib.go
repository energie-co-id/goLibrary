package mqttLib

import (
	"fmt"
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
	fmt.Println("Connected")
}

var connectLostHandler mqtt.ConnectionLostHandler = func(client mqtt.Client, err error) {
	fmt.Printf("Connect lost: %v", err)
}

func publish(client mqtt.Client, publishMessage string, publishTopic string) (string, error) {
	// client.Publish(publishTopic, 0, false, publishMessage)
	token := client.Publish(publishTopic, 0, false, publishMessage)
	token.Wait()
	// time.Sleep(time.Second)

	return publishMessage, nil
}
func subscribe(client mqtt.Client, topic string) {
	token := client.Subscribe(topic, 1, nil)
	token.Wait()
	fmt.Printf("Subscribed to topic %s", topic)
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
	publishResult, err := publish(client, publishMessage, publishTopic)
	if err != nil || publishResult != publishMessage {
		return "", err
	}
	client.Disconnect(250)

	return publishResult, err
}

func PublishAndListening(publishMessage string, publishTopic string, listeningTopic string) string {
	//config mqtt
	client := mqttConfig()
	subscribe(client, listeningTopic)
	publish(client, publishMessage, publishTopic)

	timer := 10 * 1000                                        //10 seconds in microseconds
	for receivedMessage == "no response" && timer < 10*1000 { //wait 10 seconds for the message in range of miliseconds
		time.Sleep(time.Millisecond)
		timer++
	}

	client.Disconnect(250)

	return receivedMessage
}
