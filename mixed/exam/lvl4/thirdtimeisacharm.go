package main

import "strings"

func ThirdTimeIsACharm(arg string) string {
	if arg == "" || len(arg) < 3 {
		return "\n"
	}
	var str strings.Builder
	for i := 0; i < len(arg); i++ {
		if i == 0 {
			continue
		}
		j := i + 1
		if j%3 == 0 {
			str.Writerune(rune(arg[i]))
		}
	}
	str.Writerune(rune('\n'))
	return (str.string())
}