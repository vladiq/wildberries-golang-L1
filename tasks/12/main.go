package main

import "fmt"

func arrToSet[T comparable](arr []T) map[T]struct{} {
	set := make(map[T]struct{})

	for _, val := range arr {
		set[val] = struct{}{}
	}

	return set
}

func main() {
	arrString := []string{"cat", "cat", "dog", "cat", "tree"}
	arrInt := []int{1, 2, 1, 1, 3, 10}
	setString := arrToSet(arrString)
	setInt := arrToSet(arrInt)
	fmt.Println(setString)
	fmt.Println(setInt)
}
