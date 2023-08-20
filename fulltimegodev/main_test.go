package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestState(t *testing.T) {
	state := &State{}
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			state.setState(i + 1)
			wg.Done()
		}(i)
	}

	wg.Wait()
	fmt.Println(state)
}
