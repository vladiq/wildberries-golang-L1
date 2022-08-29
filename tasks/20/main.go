package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	// разбиваем предложение на список слов и ревёрсим его
	a := strings.Split(s, " ")
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	return strings.Join(a, " ")
}

func main() {
	s := "snow dog sun"
	reversedWords := reverseWords(s)
	fmt.Println(reversedWords)
}
