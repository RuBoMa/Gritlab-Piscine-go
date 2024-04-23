package piscine

func Index(s string, toFind string) int {
	word := []rune(s)
	text := []rune(toFind)
	for countword := 0; countword <= len(word)-len(text); countword++ {
		found := true
		for countoftext := 0; countoftext < len(text); countoftext++ {
			if word[countword+countoftext] != text[countoftext] {
				found = false
				break
			}
		}
		if found {
			return countword
		}
	}
	return -1
}
