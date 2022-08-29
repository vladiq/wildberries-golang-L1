package main

import (
	"fmt"
)

func groupByTens(temperatures []float32) map[int][]float32 {
	groups := make(map[int][]float32)
	for _, t := range temperatures {
		idx := int(t) - int(t)%10
		groups[idx] = append(groups[idx], t)
	}
	return groups
}

func main() {
	temperatures := []float32{
		-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5,
	}
	fmt.Println(groupByTens(temperatures))
}
