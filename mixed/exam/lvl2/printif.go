package main

//	func main() {
//		fmt.Print(PrintIf("abcdefz"))
//		fmt.Print(PrintIf("abc"))
//		fmt.Print(PrintIf(""))
//		fmt.Print(PrintIf("14"))
//	}
func PrintIf(arg string) string {
	if len(arg) != 0 && len(arg) < 3 {
		return "Invalid Input\n"
	}
	return "G\n"
}
