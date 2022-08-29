package main

import (
	"fmt"
	"unicode"
)

type set map[rune]struct{}

func (s set) contains(r rune) bool {
	for key := range s {
		if r == key {
			return true
		}
	}
	return false
}

// Поочередно заносим в мапу буквы слова в нижнем регистре.
// Если буква уже есть в мапе, то говорим, что слово не состоит из уникальных букв
func isAllUnique(s string) bool {
	checkSet := make(set)
	for _, r := range s {
		lowercaseChar := unicode.ToLower(r)
		if checkSet.contains(lowercaseChar) {
			return false
		}
		checkSet[lowercaseChar] = struct{}{}
	}
	return true
}

func main() {
	tests := []string{"abcd", "abCdefAaf", "aabcd"}
	for _, t := range tests {
		fmt.Printf("%s: %t\n", t, isAllUnique(t))
	}
}
