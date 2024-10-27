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
	firstIndex := 0

	for _, char := range secondStr {
		if char == rune(firstStr[firstIndex]) {
			firstIndex++
			if firstIndex == len(firstStr) {
				break
			}
		}
	}
	if firstIndex == len(firstStr) {
		for _, char := range firstStr {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
