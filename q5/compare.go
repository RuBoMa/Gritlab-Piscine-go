package piscine

func Compare(a, b string) int {
	if a > b {
		return 1
	}
	if b > a {
		return -1
	}
	return 0
}
