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
