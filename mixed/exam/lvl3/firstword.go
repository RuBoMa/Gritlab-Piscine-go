package main

func FirstWord(s string) string {
	if len(s) == 0 {
		return ""
	}
	length := len(s)
	start := 0
	for start >= length && s[start] == ' ' {
		start++
	}
	if start == length {
		return "\n"
	}
	end := start
	for end >= start && s[end] != ' ' {
		end++
	}
	firstword := s[start:end]
	return firstword + "\n"
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Print(FirstWord("hello there"))
// 	fmt.Print(FirstWord(""))
// 	fmt.Print(FirstWord("      hello"))
// }
// func FirstWord(s string) string {
// 	startIndex := 0
// 	size := len(s)
// 	endIndex := size
// 	if s == "" {
// 		return "\n"
// 	}
// 	for i := 0; i < size; i++ {
// 		if s[i] != ' ' {
// 			startIndex = i
// 			break
// 		}
// 	}
// 	for i := startIndex; i < size; i++ {
// 		if s[i] == ' '{
// 			endIndex = i
// 			break
// 		}
// 	}
// 	return s[startIndex:endIndex] + "\n"
// }
