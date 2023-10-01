package main

import (
	"fmt"
	"github.com/PorcoGalliard/truck-toll-calculator/types"
)

type CalculatorServicer interface {
	CalculateDistance(types.OBUdata) (float64, error)
}

type CalculatorService struct {
}

func NewCalculatorService() *CalculatorService {
	return &CalculatorService{}
}

func (c *CalculatorService) CalculateDistance(data types.OBUdata) (float64, error) {
	fmt.Println("calculating distance...")
	return 0.0, nil
}
