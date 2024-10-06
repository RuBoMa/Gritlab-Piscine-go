package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}
func RepeatAlpha(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := ""
	for _, char := range s {
		if char >= 'a' && char <= 'z' {
			index := int(char - 'a' + 1)
			for i := 0; i < index; i++ {
				result += string(char)
			}

		} else if char >= 'A' && char <= 'Z' {
			index := int(char - 'A' + 1)
			for i := 0; i < index; i++ {
				result += string(char)
			}
		}
	}
	return result
}
