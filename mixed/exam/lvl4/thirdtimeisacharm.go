package main

func ThirdTimeIsACharm(str string) string {
	if len(str) == 0 {
		return "\n"
	}
	if len(str) < 3 {
		return "\n"
	}
	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}
	return result + "\n"
}
