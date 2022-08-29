package main

import (
	"errors"
	"fmt"
	"strconv"
)

func setIthBit(i int64, value int, position uint32) (int64, error) {
	var mask int64 = 1 << position
	switch value {
	case 0:
		return i & (^mask), nil
	case 1:
		return i | mask, nil
	default:
		return i, errors.New("the bit value should be either 0 or 1")
	}
}

func main() {
	var i int64 = 10

	res, err := setIthBit(i, 1, 0) // starts from pos 0, right to left
	if err != nil {
		panic(err)
	} else {
		fmt.Println("before:", strconv.FormatInt(i, 2))
		fmt.Println("after:", strconv.FormatInt(res, 2))
	}
}
