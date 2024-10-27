package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}

	firstStr := os.Args[1]
	secondStr := os.Args[2]

	charInSecond := make(map[rune]bool)

	for _, char := range secondStr {
		charInSecond[char] = true
	}
	seenInSecond := make(map[rune]bool)
	for _, char := range firstStr {
		if charInSecond[char] && !seenInSecond[char] {
			z01.PrintRune(char)
			seenInSecond[char] = true
		}
	}
	z01.PrintRune('\n')
}
