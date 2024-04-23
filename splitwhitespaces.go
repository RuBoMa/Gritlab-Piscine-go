package piscine

func SplitWhiteSpaces(s string) []string {
	var words []string

	start := 0
	for i, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if start < i {
				words = append(words, s[start:i])
			}
			start = i + 1
		}
	}
	if start < len(s) {
		words = append(words, s[start:])
	}
	return words
}
