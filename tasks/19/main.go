package main

import (
	"fmt"
	"unicode/utf8"
)

func reverseString(s string) string {
	a := []rune(s)
	for l, r := 0, len(a)-1; l < r; l, r = l+1, r-1 {
		a[l], a[r] = a[r], a[l]
	}
	return string(a)
}

func main() {
	testStrings := []string{
		"главрыба",
		"⨯≯⢈⓷⬸‌Ⲝ⚺ⶋ➜⠈⼒␀⏵⚛Ⰳ⇄⓶⺵⦧⺣⒗⎐⤕⪤ⶢ⋅⇗ⓗℴ❻⏣⻭⁞⊭▥ⰺ⁈∐‽♓⓲⤻ⓛⲉ⩉ⴀ⪲₦⟻⑭⧎ⴵ☥⾶⼣⹇⎭℟╪⬏⪐⩓",
	}
	for _, s := range testStrings {
		fmt.Printf("original: %s\nreversed: %s\n", s, reverseString(s))
		fmt.Println(utf8.RuneCountInString(s))
	}
}
