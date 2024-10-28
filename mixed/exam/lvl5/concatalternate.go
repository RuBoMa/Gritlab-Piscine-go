package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	var result []int
	first, second := slice1, slice2
	if len(slice1) < len(slice2) {
		first, second = slice2, slice1
	}
	for i := 0; i < len(first); i++ {
		result = append(result, first[i])
		if i < len(second) {
			result = append(result, second[i])
		}
	}
	return result
}
