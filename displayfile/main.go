package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("quest8.txt")
	if err != nil {
		fmt.Printf("the mistake is : %v\n", err.Error())
	}

	arr := make([]byte, 14)

	file.Read(arr)

	file.Close()

	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
	} else if len(os.Args) == 1 {
		fmt.Println("File name missing")
	} else {
		fmt.Println(string(arr))
	}
}
