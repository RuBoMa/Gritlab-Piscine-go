package main

import (
	"fmt"
)

func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
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
