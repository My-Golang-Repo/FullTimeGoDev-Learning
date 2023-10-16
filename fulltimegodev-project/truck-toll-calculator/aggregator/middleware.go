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

func (l *LoggingMiddleware) CalculateInvoice(obuID int) (inv *types.Invoice, err error) {
	defer func(start time.Time) {
		var (
			distance float64
			amount   float64
		)

		if inv != nil {
			distance = inv.TotalDistance
			amount = inv.TotalAmount
		}
		logrus.WithFields(logrus.Fields{
			"took":          time.Since(start),
			"err":           err,
			"obuID":         obuID,
			"totalDistance": distance,
			"totalAmount":   amount,
		}).Info("CalculateInvoice")
	}(time.Now())
	inv, err = l.next.CalculateInvoice(obuID)
	return
}
