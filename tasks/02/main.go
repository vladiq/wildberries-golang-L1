package main

import (
	"fmt"
	"sync"
)

func takeSquareUnordered(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num * num)
}

func takeSquareOrdered(ch <-chan int) {
	num := <-ch
	fmt.Println(num * num)
}

func main() {
	input := []int{2, 4, 6, 8, 10}

	fmt.Println("Unordered output using wg.WaitGroup")
	var wg sync.WaitGroup
	for _, i := range input {
		wg.Add(1)
		go takeSquareUnordered(i, &wg)
	}
	wg.Wait()

	fmt.Println("Ordered output using channels")
	ch := make(chan int)
	for _, i := range input {
		go takeSquareOrdered(ch)
		ch <- i
	}
}
