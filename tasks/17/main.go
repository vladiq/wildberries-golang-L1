package main

import (
	"errors"
	"fmt"
	"golang.org/x/exp/constraints"
	"strconv"
)

func binarySearch[T constraints.Ordered](arr []T, val T) (int, error) {
	var (
		l = 0
		r = len(arr) - 1
		m int
	)
	for l <= r {
		m = (l + r) / 2
		switch {
		case val < arr[m]:
			r = m - 1
		case val > arr[m]:
			l = m + 1
		default:
			return m, nil
		}
	}
	return -1, errors.New("value not found in the array")
}

func main() {
	// протестируем для строк и интов
	arrInt := []int{1, 10, 15, 20, 100}

	var arrString []string
	for _, i := range arrInt {
		arrString = append(arrString, strconv.Itoa(i))
	}

	if i, err := binarySearch(arrInt, 15); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}

	if i, err := binarySearch(arrString, "15"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(i)
	}
}
