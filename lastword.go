package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
func LastWord(s string) string {
	words := strings.Fields(s)
	if len(words) > 0 {
		return words[len(words)-1] + "\n"
	}
	return "\n"
}
