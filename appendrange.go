package piscine

func AppendRange(min, max int) []int {
	var vastaus []int
	if min >= max {
		return nil
	}
	for i := 0; i < max-min; i++ {
		vastaus = append(vastaus, min+i)
	}
	return vastaus
}
