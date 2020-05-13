package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/domain/model"
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/infrastructure/pubsub"
	"github.com/MegaBlackLabel/go-clean-architecture-mqtt-practice/registry"
)

func onMessageReceived(message model.Message) {
	fmt.Printf("[MQTT] Received to %s [Received Message: %s]\n", message.Topic, message.Payload)
}

func main() {
	hostname, _ := os.Hostname()
	//
	config := pubsub.MqttConfig{
		Server:    "tcp://localhost:1883",
		Sendtopic: "MQTT/Client/Update/TEST",
		Resvtopic: "MQTT/+/Update/#",
		Qos:       1,
		Retained:  false,
		Clientid:  hostname + strconv.Itoa(time.Now().Second()),
		Username:  "",
		Password:  "",
	}

	//
	mqtt, err := registry.InitialiseMqtt(config)
	if err == nil {
		mqtt.Subscribe(onMessageReceived)
		mqtt.Publish("test")
		fmt.Printf("mqtt %v", mqtt)

		for {
			time.Sleep(5000 * time.Millisecond)
			if err := mqtt.Publish("test massage"); err != nil {
				fmt.Println(err)
				os.Exit(2)
			}
		}
	}
}
