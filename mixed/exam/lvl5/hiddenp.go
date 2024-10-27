package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 3 {
		return
	}
	s1 := os.Args[1]
	s2 := os.Args[2]

	hiddenIndex := 0
	mainIndex := 0

	for mainIndex < len(s2) {
		if s1[hiddenIndex] == s2[mainIndex] {
			hiddenIndex++
			if hiddenIndex == len(s1) {
				z01.PrintRune('1')
				z01.PrintRune('\n')
				return
			}
		}
		mainIndex++
	}
	z01.PrintRune('0')
	z01.PrintRune('\n')
}
