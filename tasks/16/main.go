package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, 100)
	for i := range arr {
		arr[i] = rand.Intn(301)
	}
	fmt.Println(arr)
	fmt.Println()

	quickSort(arr, 0, len(arr)-1)
	fmt.Println(arr)
}

func quickSort(arr []int, low, high int) {
	if low >= high {
		return
	}
	pivot := partition(arr, low, high)

	quickSort(arr, low, pivot-1)
	quickSort(arr, pivot+1, high)
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}
