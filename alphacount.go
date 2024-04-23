package piscine

func AlphaCount(s string) int {
	word := []rune(s)
	count := 0
	for i := 0; i < len(word); i++ {
		if word[i] >= 97 && word[i] <= 122 || word[i] >= 65 && word[i] <= 90 {
			count++
		}
	}
	return count
}
