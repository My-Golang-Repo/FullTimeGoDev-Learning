package main

import (
	"encoding/json"
	"fmt"
	"github.com/PorcoGalliard/truck-toll-calculator/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

var neuKafkaTopic = "obudata"

type DataProducer interface {
	ProduceData(types.OBUdata) error
}

type KafkaProducer struct {
	producer *kafka.Producer
}

func NewKafkaProducer() (*KafkaProducer, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		return nil, err
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &KafkaProducer{
		producer: p,
	}, nil
}

func (p *KafkaProducer) ProduceData(data types.OBUdata) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	return p.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &neuKafkaTopic, Partition: kafka.PartitionAny},
		Value:          b,
	}, nil)
}
