package piscine

func CountChar(str string, c rune) int {
	if len(str) == 0 {
		return 0
	}
	count := 0
	for _, p := range str {
		if p == c {
			count++
		}
	}
	return count
}
