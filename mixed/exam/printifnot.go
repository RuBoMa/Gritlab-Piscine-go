package main

func printifnot(arg string) string {
	if len(arg) <= 3 {
		return "Invalid Input\n"
	} else {
		return "G\n"
	}
}
