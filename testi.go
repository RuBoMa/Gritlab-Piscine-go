package main

import (
	"fmt"
	"strconv"
	"strings"
)

var punctuationMarks = map[string]bool{
	",": true,
	".": true,
	"!": true,
	"?": true,
	":": true,
	";": true,
}

func modifySlice(s string) []string {
	slice := strings.Fields(s)
	// Define a map for vowels marks
	vowels := map[string]bool{
		"a": true,
		"e": true,
		"i": true,
		"o": true,
		"u": true,
		"A": true,
		"E": true,
		"I": true,
		"O": true,
		"U": true,
	}
	// if for loop encounters vowel and the next word starts with a vowel it replaces vowel with "an"
	for i := 0; i < len(slice); i++ {
		word := slice[i]
		if (word == "a" || word == "A") && i+1 < len(slice) && vowels[string(slice[i+1][0])] {
			slice[i] = "an"
		}
	}
	return slice
}
func modifytext(slice []string) string {

	// fmt.Println(slice)

	flag := false
	for i := 0; i < len(slice); i++ {
		word := []byte(slice[i])
		if slice[i] == "(hex)" && i > 0 {
			decimalValue := convertHexToDecimal(slice[i-1])
			slice[i-1] = decimalValue
			slice[i] = ""
		} else if slice[i] == "(bin)" && i > 0 {
			decimalValue := convertBinToDecimal(slice[i-1])
			slice[i-1] = decimalValue
			slice[i] = ""
		} else if slice[i] == "(up)" && i > 0 {
			slice[i-1] = strings.ToUpper(slice[i-1])
			slice[i] = ""
		} else if slice[i] == "(low)" && i > 0 {
			slice[i-1] = strings.ToLower(slice[i-1])
			slice[i] = ""
		} else if slice[i] == "(cap)" && i > 0 {
			slice[i-1] = strings.Title(strings.ToLower(slice[i-1]))
			slice = append(slice[:i], slice[i+1:]...)
			i--

		} else if (slice[i] == "(up,") || (slice[i] == "(low,") || (slice[i] == "(cap,") {
			num := extractNumber(slice[i+1])
			if num > 0 && i >= num {
				for j := 1; j <= num; j++ {
					switch {
					case (slice[i] == "(up,"):
						slice[i-j] = strings.ToUpper(slice[i-j])
					case (slice[i] == "(low,"):
						slice[i-j] = strings.ToLower(slice[i-j])
					case (slice[i] == "(cap,"):
						slice[i-j] = strings.Title(strings.ToLower(slice[i-j]))
					}
				}
			}
			slice = append(slice[:i], slice[i+2:]...)
			i--
		} else if punctuationMarks[slice[i]] && i > 0 {
			slice[i-1] += slice[i]
			slice = append(slice[:i], slice[i+1:]...)
			i--

			// Handle punctuation attached to words
		} else if len(word) > 1 && punctuationMarks[string(word[0])] {

			j := 0
			for j < len(word) && punctuationMarks[string(word[j])] {
				j++
			}
			if i > 0 {
				slice[i-1] += string(word[:j])
			}
			slice[i] = " " + string(word[j:])

			// Handle single quotes
		} else if slice[i] == "'" {
			if !flag {
				slice[i+1] = "'" + slice[i+1]
				slice = append(slice[:i], slice[i+1:]...)
				i--
				flag = true
			} else {
				slice[i-1] += "'"
				slice = append(slice[:i], slice[i+1:]...)
				i--
				flag = false
			}
		}
	}
	return strings.Join(slice, " ")
}

func extractNumber(s string) int {
	end := strings.Index(s, ")")
	if end < 0 {
		end += len(s)
	}
	if end > 0 {
		number, err := strconv.Atoi(s[:end])
		if err == nil {
			return number
		}
	}
	return 1 // Default to 1 if no number is found or there's an error
}
func convertHexToDecimal(hexStr string) string {
	num, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("Invalid input, (hex)")
		return hexStr
	}
	return strconv.FormatInt(num, 10)
}

func convertBinToDecimal(binStr string) string {
	num, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		fmt.Println("Invalid input, (bin)")
		return binStr
	}
	return strconv.FormatInt(num, 10)
}

func main() {

	text := "' I have, to pack 101 (bin) outfits .  ......  Packed 1a (hex) just to be sure '"

	modifedSlice := modifySlice(text)

	modifiedText := modifytext(modifedSlice)

	modifiedText = strings.Replace(modifiedText, "  ", " ", -1)

	fmt.Println(modifiedText)

}
