package main

import (
	"fmt"
	"sync"
)

type State struct {
	mu    sync.Mutex
	count int
}

func (s *State) setState(i int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count = i
}

func main() {
	state := &State{}

	for i := 0; i < 10; i++ {
		state.count = i
	}

	fmt.Println(state)
}
