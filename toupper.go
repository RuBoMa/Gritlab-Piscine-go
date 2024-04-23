package piscine

func ToUpper(s string) string {
	sentence := []rune(s)
	for i := 0; i < len(sentence); i++ {
		if s[i] >= 97 && s[i] <= 122 {
			sentence[i] -= 32
		}
	}
	return string(sentence)
}
