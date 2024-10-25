package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
	}
	input := os.Args[1]
	inWord := false

	for i := 0; i < len(input); i++ {
		if input[i] != ' ' && input[i] != '\t' {
			if inWord {
				z01.PrintRune(' ')
			}
			for i < len(input) && input[i] != ' ' && input[i] != '\t' {
				z01.PrintRune(rune(input[i]))
				i++
			}
			inWord = true
		}
	}
	z01.PrintRune('\n')
}
