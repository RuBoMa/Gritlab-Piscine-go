package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func printInt(n int) {
	if n < 0 {
		z01.PrintRune('-')
		n = -n
	}
	if n >= 10 {
		printInt(n / 10)
	}
	z01.PrintRune(rune(n%10) + '0')
}

func main() {
	point := &point{}
	setPoint(point)
	z01.PrintRune('x')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printInt(point.x)
	z01.PrintRune(',')
	z01.PrintRune(' ')
	z01.PrintRune('y')
	z01.PrintRune(' ')
	z01.PrintRune('=')
	z01.PrintRune(' ')
	printInt(point.y)
	z01.PrintRune('\n')
}
