package solutions

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}
	var result []byte
	for i := 0; i < len(s); i++ {
		c := s[i]

		if (c < 'a' || c > 'z') && (c < 'A' || c > 'Z') {
			return s
		}
		if c >= 'A' && c <= 'Z' {
			if i > 0 {
				if s[i-1] >= 'A' && s[i-1] <= 'Z' || i == len(s)-1 {
					return s
				}
				result = append(result, '_')
			}
			c += 32
		}
		result = append(result, c)
	}
	return string(result)
}
