package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	vastaus := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		vastaus[i] = min + i
	}
	return vastaus
}
