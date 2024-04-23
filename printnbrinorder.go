package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	digits := []int{}

	for n > 0 {
		digit := n % 10
		digits = append(digits, digit)
		n /= 10
	}
	for i := 0; i <= 9; i++ {
		for _, digit := range digits {
			if digit == i {
				z01.PrintRune(rune('0' + i))
			}
		}
	}
}
