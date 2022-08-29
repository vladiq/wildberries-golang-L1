package main

import (
	"fmt"
	"sync"
)

type sumOfSquares struct {
	mu  sync.Mutex
	sum int
}

func (s *sumOfSquares) AddSquare(wg *sync.WaitGroup, val int) {
	defer wg.Done()

	s.mu.Lock()
	s.sum += val * val
	s.mu.Unlock()
}

func main() {
	arr := []int{2, 4, 6, 8, 10}

	var wg sync.WaitGroup

	s := sumOfSquares{}
	for _, i := range arr {
		wg.Add(1)
		go s.AddSquare(&wg, i)
	}

	wg.Wait()
	fmt.Println(s.sum)
}
