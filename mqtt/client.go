package mqtt

import (
	"fmt"
	"net"

	proto "github.com/huin/mqtt"
	"github.com/jeffallen/mqtt"
)

//Client Subscripton struct
type Client struct {
	Port string
}

//NewClient declares a new broker
func NewClient(port string) *Client {
	m := &Client{Port: port}
	fmt.Println(m)
	return m
}

//String returns details of the MqttClient
func (c *Client) String() string {
	return fmt.Sprintf("MqttClient: IP %s", c.Port)
}

//Publish MQTT message, takes topic as a string, data as a byte array and retain flag as bool
func (c *Client) Publish(topic string, data string, retain bool) error {

	con, err := net.Dial("tcp", c.Port)
	if err != nil {
		return err
	}

	ccPub := mqtt.NewClientConn(con)

	err = ccPub.Connect("", "")
	if err != nil {
		return err
	}

	ccPub.Publish(&proto.Publish{
		Header: proto.Header{
			Retain: retain,
		},
		TopicName: topic,
		Payload:   proto.BytesPayload(data),
	})

	ccPub.Disconnect()

	return nil
}
