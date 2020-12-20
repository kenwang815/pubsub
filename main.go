package main

import (
	"fmt"
	"math/rand"
	"time"

	"example.pubsub/broker"
	"example.pubsub/sensor"
)

func publisTo(eb broker.IEventBus, topic string, sensorName string) {
	s := sensor.NewSensor(sensorName)
	for {
		eb.Publish(topic, s.GetData())
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}

func watchData(ch string, dataEvent broker.DataEvent) {
	fmt.Printf("Channel: %s, Topic: %s, DataEvent: %v\n", ch, dataEvent.Topic, dataEvent.Data)
}

func main() {
	eb := broker.NewEventBus()

	sensorChannel := make(chan broker.DataEvent)
	eb.Subscribe("sensor", sensorChannel)

	go publisTo(eb, "sensor", "bedroom")

	for {
		select {
		case d := <-sensorChannel:
			go watchData("sensorChannel", d)
		}
	}
}
