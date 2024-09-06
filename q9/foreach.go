package piscine

var output string

func ForEach(f func(int), a []int) {
	for _, num := range a {
		f(num)
	}
}

func PrintNbr(num int) {
	if num == 0 {
		Str("0")
		return
	}
	var digits [20]byte
	i := 0
	negative := false

	if num < 0 {
		negative = true
		num = -num
	}
	for num > 0 {
		digits[i] = byte(num%10) + '0'
		num /= 10
		i++
	}
	if negative {
		Str("-")
	}
	for j := i - 1; j >= 0; j-- {
		Pune(rune(digits[j]))
	}
	FlushOutput()
}

func Str(s string) {
	output += s
}

func Pune(r rune) {
	output += string(r)
}

func FlushOutput() {
	for _, char := range output {
		Pune(char)
	}
	output = ""
}
