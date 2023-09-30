package main

import "github.com/PorcoGalliard/truck-toll-calculator/types"

type DataConsumer interface {
	ConsumeData(data types.OBUdata) error
}
