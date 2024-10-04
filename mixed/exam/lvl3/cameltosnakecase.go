package solutions

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}
	result := ""
	for i, char := range s {
		if (char < 'a' || char > 'z') && (char < 'A' || char > 'Z') {
			return s
		}
		if char >= 'A' && char <= 'Z' {
			if i > 0 {
				if s[i-1] <= 'A' && s[i-1] >= 'Z' || i == len(s)-1 {
					return s
				}
				result += "_"
			}
		}
		result += string(char)
	}
	return result
}
