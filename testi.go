package main

import (
	"fmt"
	"strconv"
	"strings"
)

func modifySlice(slice []string) []string {
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
func modifytext(s string) string {
	slice := strings.Fields(s)
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
		} else if slice[i] == "," || slice[i] == "." || slice[i] == "!" || slice[i] == "?" || slice[i] == ":" || slice[i] == ";" || slice[i] == "..." || slice[i] == "!?" {
			slice[i-1] += slice[i]
			slice = append(slice[:i], slice[i+1:]...)
			i--

		} else if word[0] == ',' || word[0] == '.' || word[0] == '!' || word[0] == '?' || word[0] == ':' || word[0] == ';' {
			if word[0] == '.' && word[1] == '.' && word[2] == '.' {
				slice[i] = string(word[3:])
				slice[i-1] += string(word[0:3])
			} else if word[0] == '!' && word[1] == '?' {
				slice[i] = string(word[2:])
				slice[i-1] += string(word[0:2])
			} else {
				slice[i] = string(word[1:])
				slice[i-1] += string(word[0])
			}
		} else if slice[i] == "'" {
			if !flag {
				var word_after []byte
				word_after = append(word_after, '\'')
				word_after = append(word_after, []byte(slice[i+1])...)
				slice[i+1] = string(word_after)
				slice = append(slice[:i], slice[i+1:]...)
				i--
				flag = true
			} else {
				slice[i-1] += "'"
				slice = append(slice[:i], slice[i+1:]...)
				i-- // Adjust index after removing the current element
				flag = false
			}
		}
	}
	return strings.Join(slice, " ")
}

func extractNumber(s string) int {
	//start := strings.Index(s, ",") + 1
	end := strings.Index(s, ")")
	if end < 0 {
		end += len(s)
	}
	// fmt.Println(s[:end])
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
		return hexStr
	}
	return strconv.FormatInt(num, 10)
}

func convertBinToDecimal(binStr string) string {
	num, err := strconv.ParseInt(binStr, 2, 64)
	if err != nil {
		return binStr
	}
	return strconv.FormatInt(num, 10)
}

func main() {

	text := "I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure"

	// modifiedText := modifySlice(strings.Fields(text))
	finaltext := modifytext(text)

	fmt.Println(finaltext)

}
