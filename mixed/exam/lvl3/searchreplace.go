package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) > 4 {
		return
	}
	word := os.Args[1]
	new := os.Args[2]
	old := os.Args[3]

	result := ""
	for _, i := range word {
		if string(i) == old {
			result += new
		} else {
			result += string(i)
		}
	}
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
