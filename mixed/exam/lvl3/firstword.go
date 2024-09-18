package main

import "strings"

func FirstWord(s string) string {
	words := strings.Fields(s)

	if len(s) > 0 {
		return words[0] + "\n"
	}
	return "\n"
}
