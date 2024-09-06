package piscine

func Unmatch(a []int) int {
	for _, num := range a {
		count := 0
		for _, n := range a {
			if num == n {
				count++
			}
		}
		if count%2 != 0 {
			return num
		}
	}
	return -1
}
