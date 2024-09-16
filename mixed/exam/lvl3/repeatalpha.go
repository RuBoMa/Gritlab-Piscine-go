package main

import (
	"unicode"
)

//	func main() {
//		fmt.Println(RepeatAlpha("abc"))
//		fmt.Println(RepeatAlpha("Choumi."))
//		fmt.Println(RepeatAlpha(""))
//		fmt.Println(RepeatAlpha("abacadaba 01!"))
//	}
func RepeatAlpha(s string) string {
	res := ""
	for _, r := range s {
		if unicode.IsLetter(r) {
			rep := unicode.ToLower(r) - 'a' + 1
			for i := 0; i < int(rep); i++ {
				res += string(r)
			}
		} else {
			res += string(r)
		}
	}
	return res
}
