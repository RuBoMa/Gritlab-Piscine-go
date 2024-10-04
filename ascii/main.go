package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		return
	}
	input := os.Args[1]
	fmt.Println(PrintAsciiArt(input))
}

func PrintAsciiArt(input string) string {
	input = strings.ReplaceAll(input, "\\n", "\n")
	inputfile, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Print("error", err)
		return ""
	}

	inputFileLines := strings.Split(string(inputfile), "\n")
	words := strings.Split(input, "\n")
	result := ""

	for _, word := range words {
		for i := 0; i < 8; i++ {
			for _, char := range word {
				result += inputFileLines[i+(int(char-' ')*9)+1]
			}
			result += "\n"
		}
	}
	return result
}
