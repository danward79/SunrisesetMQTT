package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/danward79/sunrise"
	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"
)

//MqttClient Subscripton struct
type MqttClient struct {
	Port string
}

//newClient declares a new broker
func newClient(port string) *MqttClient {
	m := &MqttClient{Port: port}
	fmt.Println(m)
	return m
}

//String returns details of the MqttClient
func (c *MqttClient) String() string {
	return fmt.Sprintf("MqttClient: IP %s", c.Port)
}

//Command line switches
var lat = flag.Float64("y", -37.81, "Enter the Latitude of your location")
var long = flag.Float64("x", 144.96, "Enter the Longitude of your location")
var mqttServer = flag.String("s", ":1883", "Enter the IP and Port of your MQTT Broker. e.g. 127.0.0.1:1883")
var retain = flag.Bool("r", false, "Use r to retain the last state broadcast with MQTT")

func main() {
	flag.Parse()

	mqttClient := newClient(*mqttServer)
	ch := sunrise.New(*lat, *long).Start()

	for {
		select {
		case m := <-ch:
			topic := "home/" + m["location"] + "/state"
			mqttClient.publish(topic, m["state"], *retain)
		}
	}
}

//publish MQTT message, takes topic as a string, data as a byte array and retain flag as bool
func (c *MqttClient) publish(topic string, data string, retain bool) {

	con, err := net.Dial("tcp", c.Port)
	gotError(err)

	ccPub := mqtt.NewClientConn(con)

	err = ccPub.Connect("", "")
	gotError(err)

	ccPub.Publish(&proto.Publish{
		Header: proto.Header{
			Retain: retain,
		},
		TopicName: topic,
		Payload:   proto.BytesPayload(data),
	})

	ccPub.Disconnect()
}

//Generic Function to catch errors
func gotError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
