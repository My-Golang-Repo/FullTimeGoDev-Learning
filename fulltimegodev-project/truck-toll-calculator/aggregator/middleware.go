package main

import (
	"github.com/PorcoGalliard/truck-toll-calculator/types"
	"github.com/sirupsen/logrus"
	"time"
)

type LoggingMiddleware struct {
	next Aggregator
}

func NewLogMiddleware(next Aggregator) Aggregator {
	return &LoggingMiddleware{
		next: next,
	}
}

func (l *LoggingMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
			"func": "AggregateDistance",
		}).Info("Aggregate Distance")
	}(time.Now())
	err = l.next.AggregateDistance(distance)
	return
}
