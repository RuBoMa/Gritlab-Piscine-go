package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]

	if !checkInput(arguments) {
		return
	}

	var table [9][9]rune
	table = fillTable(table, arguments)

	if !isSolve(&table) {
		fmt.Println("Error")
		return
	}

	printTable(table)
}

func checkInput(args []string) bool {
	if len(args) != 9 {
		fmt.Println("Error") // input length is out of range
		return false
	}

	for i := 0; i < 9; i++ {
		if len(args[i]) != 9 {
			fmt.Println("Error")
			return false
		}

		for _, value := range args[i] {
			if (value < '1' || value > '9') && value != '.' {
				fmt.Println("Error")
				return false
			}
		}
	}
	return true
}

// fills table with input args
func fillTable(table [9][9]rune, args []string) [9][9]rune {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			table[i][j] = rune(args[i][j])
		}
	}
	return table
}

// counts remaining empty cells
func isDots(table *[9][9]rune) bool {
	for _, row := range table {
		for _, value := range row {
			if value == '.' {
				return true
			}
		}
	}
	return false
}

func isValid(table *[9][9]rune, x, y int, z rune) bool {
	for i := 0; i < 9; i++ {
		if z == table[i][x] {
			return false
		}
	}

	for j := 0; j < 9; j++ {
		if z == table[y][j] {
			return false
		}
	}

	a, b := x/3*3, y/3*3

	for i := a; i < a+3; i++ {
		for j := b; j < b+3; j++ {
			if z == table[j][i] {
				return false
			}
		}
	}
	return true
}

// fills empty cells recursively
func isSolve(table *[9][9]rune) bool {
	if !isDots(table) {
		return true
	}

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if table[y][x] == '.' {
				for z := '1'; z <= '9'; z++ {
					if isValid(table, x, y, z) {
						table[y][x] = z
						if isSolve(table) {
							return true
						}
						table[y][x] = '.'
					}
				}
				return false
			}
		}
	}
	return false
}

// prints solved sudoku table
func printTable(table [9][9]rune) {
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			if x != 8 {
				fmt.Printf("%c ", table[y][x])
			} else {
				fmt.Printf("%c", table[y][x])
			}
		}
		fmt.Println()
	}
}
