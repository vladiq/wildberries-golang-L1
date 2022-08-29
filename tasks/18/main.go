package main

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	mu      sync.Mutex
	counter int
}

func (sm *safeCounter) increment(wg *sync.WaitGroup) {
	defer wg.Done()

	sm.mu.Lock()
	defer sm.mu.Unlock()
	(*sm).counter++
}

func main() {
	sm := safeCounter{}

	var wg sync.WaitGroup
	for i := 0; i < 1e6; i++ {
		wg.Add(1)
		go sm.increment(&wg)
	}

	wg.Wait()
	fmt.Println(sm.counter)
}
