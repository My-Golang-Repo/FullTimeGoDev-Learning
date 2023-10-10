package main

import (
	"encoding/json"
	"github.com/PorcoGalliard/truck-toll-calculator/aggregator/client"
	"github.com/PorcoGalliard/truck-toll-calculator/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
	"time"
)

// This can also be called Kafka Transport
type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorServicer
	aggClient   *client.Client
}

func NewKafkaConsumer(topic string, svc CalculatorServicer, aggClient *client.Client) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		return nil, err
	}

	return &KafkaConsumer{
		consumer:    c,
		calcService: svc,
		aggClient:   aggClient,
	}, nil
}

func (c *KafkaConsumer) Start() {
	logrus.Info("kafka transport started")
	c.isRunning = true
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunning {
		// msg di bawah ini bentuknya byte guys
		msg, err := c.consumer.ReadMessage(-1)
		if err != nil {
			logrus.Errorf("Kafka consume error %s", err)
			continue
		}
		var data types.OBUdata
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			logrus.Errorf("JSON serialization error %s", err)
			continue
		}
		distance, err := c.calcService.CalculateDistance(data)
		if err != nil {
			logrus.Errorf("Calculation error %s", err)
		}
		req := types.Distance{
			Value: distance,
			Unix:  time.Now().UnixNano(),
			OBUID: data.OBUID,
		}

		if err := c.aggClient.AggregateInvoice(req); err != nil {
			logrus.Errorf("aggregate error", err)
			continue
		}
	}
}
