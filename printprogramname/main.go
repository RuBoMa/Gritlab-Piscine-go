package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[0]
	for i, char := range programName {
		if i > 1 {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}
