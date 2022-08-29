package main

import (
	"fmt"
	"time"
)

func sleep(duration time.Duration) {
	startTime := time.Now()
	// находимся в цикле, пока время, прошедшее с запуска функции меньше запрашиваемого времени
	for {
		if time.Since(startTime) >= duration {
			break
		}
	}
}

func main() {
	startTime := time.Now()
	defer func() {
		fmt.Println("Time passed:", time.Since(startTime))
	}()

	sleep(time.Second)
}
