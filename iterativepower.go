package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if power == 0 {
		return 1
	}
	if power == 1 {
		return nb
	} else {
		factorial := nb
		for i := 2; i <= power; i++ {
			factorial *= nb
		}
		return factorial
	}
}
