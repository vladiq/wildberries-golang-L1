package main

import (
	"fmt"
	"time"
)

func writeAndRead(channel chan int, timeToWork time.Duration) {
	startTime := time.Now()

	valToSend := 0
	for time.Since(startTime) < timeToWork {
		go func() {
			channel <- valToSend
		}()
		valToSend++

		fmt.Println(<-channel)
	}
}

func main() {
	myChan := make(chan int)
	writeAndRead(myChan, time.Second)
}
