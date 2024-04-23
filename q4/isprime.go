package piscine

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	} else if nb == 2 {
		return true
	}
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}
