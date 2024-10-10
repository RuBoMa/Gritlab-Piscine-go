package main

import (
	"fmt"
)

func main() {
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("      hello"))
}
func FirstWord(s string) string {
	startIndex := 0
	size := len(s)
	endIndex := size
	if s == "" {
		return "\n"
	}
	for i := 0; i < size; i++ {
		if s[i] != ' ' {
			startIndex = i
			break
		}
	}
	for i := startIndex; i < size; i++ {
		if s[i] == ' '{
			endIndex = i
			break
		}
	}
	return s[startIndex:endIndex] + "\n"
}
