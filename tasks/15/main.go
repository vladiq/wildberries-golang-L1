package main

import (
	"fmt"
	"golang.org/x/exp/rand"
	"unicode/utf8"
)

func createHugeString(length int) string {
	sample := []rune("привет(つ◉益◉)つпока")
	generatedString := make([]rune, length)
	for i := 0; i < length; i++ {
		generatedString[i] = sample[rand.Intn(len(sample))]
	}
	return string(generatedString)
}

func main() {
	var justString string // сделаем justString локальной переменной
	hugeString := createHugeString(1 << 10)

	// неправильный способ взятия слайса
	justString = hugeString[:100]
	fmt.Printf("Wrong way result: \"%s\"\nlen([]byte): %d, len([]rune): %d\n\n",
		justString, len(justString), utf8.RuneCountInString(justString))

	// правильно будет преобразовать строку в массив рун и взять слайс
	justString = string([]rune(hugeString)[:100])
	fmt.Printf("Right way result: \"%s\"\nlen([]byte): %d, len([]rune): %d\n",
		justString, len(justString), utf8.RuneCountInString(justString))
}
