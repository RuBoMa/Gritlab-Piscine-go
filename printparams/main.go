package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	programName := os.Args[1:]
	for _, char := range programName {
		for _, c := range char {
			z01.PrintRune(c)
		}
		z01.PrintRune('\n')
	}
}
