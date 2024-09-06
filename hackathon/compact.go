package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	nonZeroCount := 0
	j := 0

	for _, s := range slice {
		if s != "" {
			slice[j] = s
			j++
			nonZeroCount++
		}
	}
	*ptr = make([]string, nonZeroCount)
	copy(*ptr, slice[:j])
	return nonZeroCount
}
