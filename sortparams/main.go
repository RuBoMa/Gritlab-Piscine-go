package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	CommandArg := os.Args[1:]
	for i := 0; i < len(CommandArg)-1; i++ {
		min := i
		for j := i + 1; j < len(CommandArg); j++ {
			if CommandArg[j] < CommandArg[min] {
				min = j
			}
		}
		CommandArg[i], CommandArg[min] = CommandArg[min], CommandArg[i]
	}
	for _, arg := range CommandArg {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
