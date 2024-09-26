package main

func RetainFirstHalf(str string) string {
	half := len(str)
	if half == 0 {
		return ""
	}
	if half == 1 {
		return str
	}
	jako := half / 2
	return str[:jako]
}
