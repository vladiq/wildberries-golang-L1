package main

import "fmt"

func writeX(x int, c chan<- int) {
	c <- x
}

func writeXMult2(x int, c chan<- int) {
	c <- 2 * x
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	// так как мы не выводим содержимое канала chanForX,
	// то его можно было бы объявить как make(chan int, len(arr))
	chanForX := make(chan int)
	chanForXMult2 := make(chan int)

	for _, x := range arr {
		go writeX(x, chanForX)
		go writeXMult2(x, chanForXMult2)
		<-chanForX
		fmt.Println(<-chanForXMult2)
	}
}
