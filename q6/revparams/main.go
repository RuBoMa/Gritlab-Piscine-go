package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[1:]
	for i := len(programName) - 1; i >= 0; i-- {
		for _, char := range programName[i] {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
