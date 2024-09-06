package piscine

func Capitalize(s string) string {
	sentence := []rune(s)
	firstLetter := true

	for i := 0; i < len(sentence); i++ {
		if (sentence[i] >= 'A' && sentence[i] <= 'Z') || (sentence[i] >= 'a' && sentence[i] <= 'z') || (sentence[i] >= '0' && sentence[i] <= '9') {
			if firstLetter {
				if sentence[i] >= 'a' && sentence[i] <= 'z' {
					sentence[i] -= 32
				}
				firstLetter = false
			} else {
				if sentence[i] >= 'A' && sentence[i] <= 'Z' {
					sentence[i] += 32
				}
			}
		} else {
			firstLetter = true
		}
	}
	return string(sentence)
}
