package main

import (
	"github.com/PorcoGalliard/truck-toll-calculator/types"
	"github.com/sirupsen/logrus"
	"time"
)

type LoggingMiddleware struct {
	next DataProducer
}

func NewLogMiddleware(next DataProducer) *LoggingMiddleware {
	return &LoggingMiddleware{
		next: next,
	}
}

func (l *LoggingMiddleware) ProduceData(data types.OBUdata) error {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"obuID": data.OBUID,
			"long":  data.Long,
			"lat":   data.Lat,
			"took":  time.Since(start),
		}).Info("producing to kafka")
	}(time.Now())
	return l.next.ProduceData(data)
}
