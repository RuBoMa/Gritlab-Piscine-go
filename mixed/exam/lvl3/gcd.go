package piscine

// func Gcd(a, b uint) uint {
// 	for a != b {
// 		if a > b {
// 			a -= b
// 		} else {
// 			b -= a
// 		}
// 	}
// 	return a
// }

func Gcd(a, b uint) uint {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
