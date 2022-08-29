package main

import (
	"fmt"
	"sync"
)

type safeMap struct {
	mu        sync.Mutex
	container map[int]int
}

func (sm *safeMap) write(wg *sync.WaitGroup, key, value int) {
	defer wg.Done()

	sm.mu.Lock()
	defer sm.mu.Unlock()

	(*sm).container[key] = value
}

func main() {
	sm := safeMap{container: make(map[int]int)}
	var wg sync.WaitGroup

	for i := 0; i <= 3e4; i++ {
		wg.Add(1)

		go sm.write(&wg, i, i)
	}

	wg.Wait()
	fmt.Println(sm.container)
}
