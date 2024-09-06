package piscine

func NRune(s string, n int) rune {
	w := []rune(s)

	if n < 1 || n > len(w) {
		return 0
	}
	return w[n-1]
}
