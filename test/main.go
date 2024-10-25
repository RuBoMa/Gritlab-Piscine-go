package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]

	for i := 0; i < len(input); i++ {
		if input[i] != ' ' {
			z01.PrintRune(rune(input[i])) 
		} else if i > 0 && input[i-1] != ' ' {
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}