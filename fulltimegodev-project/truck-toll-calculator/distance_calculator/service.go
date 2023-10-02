package main

import (
	"github.com/PorcoGalliard/truck-toll-calculator/types"
	"math"
)

type CalculatorServicer interface {
	CalculateDistance(types.OBUdata) (float64, error)
}

type CalculatorService struct {
	prevPoint []float64
}

func NewCalculatorService() (*CalculatorService, error) {
	return &CalculatorService{}, nil
}

func (c *CalculatorService) CalculateDistance(data types.OBUdata) (float64, error) {
	distance := 0.0
	if len(c.prevPoint) > 0 {
		distance = calculateDistance(c.prevPoint[0], c.prevPoint[1], data.Lat, data.Long)
	}
	c.prevPoint = []float64{data.Lat, data.Long}
	return distance, nil
}

func calculateDistance(x1, x2, y1, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))
}
