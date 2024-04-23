package piscine

func UltimateDivMod(a *int, b *int) {
	resolve1 := *a / *b
	resolve2 := *a % *b
	*a = resolve1
	*b = resolve2
}
