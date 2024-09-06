package piscine

func ToLower(s string) string {
	sentence := []rune(s)
	for i := 0; i < len(sentence); i++ {
		if s[i] >= 65 && s[i] <= 90 {
			sentence[i] += 32
		}
	}
	return string(sentence)
}
