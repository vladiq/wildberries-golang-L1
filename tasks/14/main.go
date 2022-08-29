package main

import (
	"errors"
	"fmt"
)

func deduceType(val interface{}) (string, error) {
	switch val.(type) {
	case int:
		return fmt.Sprintf("%d is an integer", val), nil
	case string:
		return fmt.Sprintf("%s is a string", val), nil
	case bool:
		return fmt.Sprintf("%t is a boolean", val), nil
	case chan int:
		return fmt.Sprintf("%v is a channel", val), nil
	default:
		return "", errors.New("unknown type")
	}
}

func main() {
	diverseSlice := []interface{}{1, "abc", true, []string{}, make(chan int)}

	for _, val := range diverseSlice {
		if t, err := deduceType(val); err != nil {
			fmt.Printf("%v: %v\n", val, err)
		} else {
			fmt.Println(t)
		}

	}
}
