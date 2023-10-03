package main

import "github.com/PorcoGalliard/truck-toll-calculator/types"

type MemoryStore struct {
}

func (m *MemoryStore) Insert(d types.Distance) error {
	return nil
}

func NewMemoryStore() *MemoryStore {
	return nil
}
