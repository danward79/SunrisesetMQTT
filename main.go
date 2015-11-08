package main

import (
	"flag"
	"log"

	"github.com/danward79/SunrisesetMQTT/mqtt"
	"github.com/danward79/SunrisesetMQTT/sunriseset"
)

//Command line switches
var lat = flag.Float64("y", -37.81, "Enter the Latitude of your location, default is Melbourne, AU")
var long = flag.Float64("x", 144.96, "Enter the Longitude of your location default is Melbourne, AU")
var mqttServer = flag.String("s", ":1883", "Enter the IP and Port of your MQTT Broker. e.g. 127.0.0.1:1883")
var retain = flag.Bool("r", false, "Use r to retain the last state broadcast with MQTT")

func main() {
	flag.Parse()

	mqttClient := mqtt.NewClient(*mqttServer)
	ch := sunriseset.New(*lat, *long).Start()

	for m := range ch {
		topic := "home/" + m["location"] + "/state"
		if err := mqttClient.Publish(topic, m["state"], *retain); err != nil {
			log.Println(err)
		}
	}
}
