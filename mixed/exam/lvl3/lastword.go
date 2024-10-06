package main

//	func main() {
//		fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
//		fmt.Print(LastWord(" lorem,ipsum "))
//		fmt.Print(LastWord(" "))
//	}
func LastWord(s string) string {
	if s == "" {
		return "\n"
	}
	length := len(s)
	end := length - 1
	for end >= 0 && s[end] == ' ' {
		end--
	}
	if end < 0 {
		return "\n"
	}
	start := end
	for start >= 0 && s[start] != ' ' {
		start--
	}
	lastword := s[start+1 : end+1]
	return lastword + "\n"
}
