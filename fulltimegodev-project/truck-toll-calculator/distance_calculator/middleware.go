package main

import (
	"github.com/PorcoGalliard/truck-toll-calculator/types"
	"github.com/sirupsen/logrus"
	"time"
)

type LoggingMiddleware struct {
	next CalculatorServicer
}

func NewLogMiddleware(next CalculatorServicer) CalculatorServicer {
	return &LoggingMiddleware{
		next: next,
	}
}

func (m *LoggingMiddleware) CalculateDistance(data types.OBUdata) (dist float64, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"dist": dist,
			"err":  err,
		}).Info("calculate distance")
	}(time.Now())

	dist, err = m.next.CalculateDistance(data)
	return
}
