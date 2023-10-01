package main

import "log"

const kafkaTopic = "obuData"

func main() {
	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
