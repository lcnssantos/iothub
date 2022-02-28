package mqtt

import "time"

type MQTT struct {
	host string
	port string
	user string
	pass string
}

func NewMQTT(host string, port string, user string, pass string) *MQTT {
	return &MQTT{host: host, port: port, user: user, pass: pass}
}

func (m MQTT) Listen(topic string, done chan interface{}) {
	for {
		time.Sleep(1 * time.Second)
		done <- struct {
			Message string `json:"message"`
			Device  string `json:"device"`
		}{Message: "new message", Device: "192.168.1.401"}
	}
}
