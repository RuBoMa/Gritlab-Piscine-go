package piscine

import (
	"strconv"
)

func ZipString(s string) string {
	result := ""
	lenght := len(s)
	index := 0

	for index < lenght {
		count := 1
		for index+1 < lenght && s[index] == s[index+1] {
			count++
			index++
		}
		result += strconv.Itoa(count) + string(s[index])
		index++
	}

	return result
}
