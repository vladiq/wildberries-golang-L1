package main

import (
	"errors"
	"fmt"
)

func removeIth[T comparable](arr []T, i int) ([]T, error) {
	if i < 0 || i >= len(arr) {
		errMsg := fmt.Sprintf("removing %dth element: index must lie between %d and %d", i, 0, len(arr)-1)
		return arr, errors.New(errMsg)
	}
	return append(arr[:i], arr[i+1:]...), nil
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	if newArr, err := removeIth(arr, 2); err != nil {
		panic(err)
	} else {
		fmt.Println(newArr)
	}
}
