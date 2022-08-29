package main

import (
	"fmt"
	"sync"
)

type safeCounter struct {
	mu      sync.Mutex
	counter int
}

func (sc *safeCounter) increment(wg *sync.WaitGroup) {
	defer wg.Done()

	sc.mu.Lock()
	defer sc.mu.Unlock()
	(*sc).counter++
}

func main() {
	sc := safeCounter{}

	var wg sync.WaitGroup
	for i := 0; i < 1e6; i++ {
		wg.Add(1)
		go sc.increment(&wg)
	}

	wg.Wait()
	fmt.Println(sc.counter)
}
