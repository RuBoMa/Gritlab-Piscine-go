package piscine

func TrimAtoi(s string) int {
	var num int
	var sign int = 1
	var started bool

	for _, char := range s {
		if char >= '0' && char <= '9' {
			num = num*10 + int(char-'0')
			started = true
		} else if char == '-' && !started {
			sign = -1
		} else if char != 0 {
			continue
		}
	}
	return num * sign
}
