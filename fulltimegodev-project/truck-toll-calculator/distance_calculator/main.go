package main

import "log"

const kafkaTopic = "obuData"

func main() {
	var (
		err error
		svc CalculatorServicer
	)

	svc, err = NewCalculatorService()
	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc)
	if err != nil {
		log.Fatal(err)
	}
	kafkaConsumer.Start()
}
