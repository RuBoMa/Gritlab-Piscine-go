package piscine

func Rot14(s string) string {
	var result string
	for _, char := range s {
		switch {
		case char >= 'a' && char <= 'z':
			newChar := 'a' + (char-'a'+14)%26
			result += string(newChar)
		case char >= 'A' && char <= 'Z':
			newChar := 'A' + (char-'A'+14)%26
			result += string(newChar)
		default:
			result += string(char)
		}
	}
	return result
}
