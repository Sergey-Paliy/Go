package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	counts := make(map[rune]int)
	for _, letter := range strings.ToLower(text) {
		if unicode.IsLetter(letter) {
			counts[letter]++
		}
	}
	for letter, count := range counts {
		fmt.Printf("%c: %d\n", letter, count)

	}
}
