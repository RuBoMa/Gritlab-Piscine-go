package main

import (
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, fileName := range os.Args[1:] {
			file, err := os.Open(fileName)
			if err != nil {
				os.Stderr.WriteString("ERROR: " + err.Error() + "\n")
				os.Exit(1)
			}
			defer file.Close()
			buf := make([]byte, 1024)
			for {
				n, err := file.Read(buf)
				if err != nil && n == 0 {
					break
				}
				os.Stdout.Write(buf[:n])
			}
		}
	} else {
		buf := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil && n == 0 {
				break
			}
			os.Stdout.Write(buf[:n])
		}
	}
}
