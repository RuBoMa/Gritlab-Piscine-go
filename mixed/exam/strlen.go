package piscine

func StrLen(s string) int {
	sentence := []rune(s)
	count := 0
	for i := range sentence {
		i++
		count++
	}
	return count
}
