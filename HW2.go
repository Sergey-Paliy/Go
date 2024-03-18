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
	letters := 0
	text = strings.ToLower(text)
	for _, letter := range text {
		if unicode.IsLetter(letter) {
			counts[letter]++
			letters++
		}
	}
	for letter, count := range counts {
		percentage := float64(count) / float64(letters) * 100
		fmt.Printf("%c: %d-> %.2f%%\n", letter, count, percentage)

	}
}
