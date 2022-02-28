package mqtt

import (
	"encoding/json"
	"fmt"
	"log"

	mqtt2 "github.com/eclipse/paho.mqtt.golang"
)

type MQTT struct {
	host   string
	port   string
	user   string
	pass   string
	client mqtt2.Client
}

func NewMQTT(host string, port string, user string, pass string) *MQTT {
	opts := mqtt2.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%s", host, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(user)
	opts.SetPassword(pass)
	client := mqtt2.NewClient(opts)

	return &MQTT{host: host, port: port, user: user, pass: pass, client: client}
}

func (m MQTT) Connect() error {
	token := m.client.Connect()

	if token.Wait() && token.Error() != nil {
		log.Println(token.Error())
		return token.Error()
	}

	return nil
}

func (m MQTT) Listen(topic string, done chan interface{}) {
	m.client.Subscribe(topic, 0, func(c mqtt2.Client, m mqtt2.Message) {
		var data map[string]interface{}

		err := json.Unmarshal(m.Payload(), &data)

		if err != nil {
			log.Println(err)
		}

		done <- data
	})
}

func (m MQTT) Close() error {
	return nil
}
