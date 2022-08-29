package main

import "fmt"

type set map[int]struct{}

func intersect(s1 set, s2 set) set {
	intersection := make(set)

	for key := range s2 {
		if _, ok := s1[key]; ok {
			intersection[key] = struct{}{}
		}
	}

	return intersection
}

func main() {
	s1 := set{
		1: struct{}{},
		2: struct{}{},
		3: struct{}{},
		4: struct{}{},
		5: struct{}{},
	}
	s2 := set{
		3: struct{}{},
		5: struct{}{},
	}
	intersection := intersect(s1, s2)
	for key := range intersection {
		fmt.Println(key)
	}
}
