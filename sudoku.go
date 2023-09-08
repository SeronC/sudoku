package main

import (
	"fmt"
	"os"
)

func main() {
	arr := os.Args[1:]
	if checkE(arr) {
		fmt.Println("ERROR!")
	} else {
		sudoku := slicer(arr)
		if solver(&sudoku, len(sudoku)) {
			printSudoku(sudoku)
		} else {
			fmt.Println("ERROR!")
		}
	}
}

func solver(arr *[][]int, len int) bool {
	empty := true
	row := -1
	col := -1
	for x := 0; x < len; x++ {
		for y := 0; y < len; y++ {
			if (*arr)[x][y] == 0 {
				row = x
				col = y
				empty = false
				break
			}
		}
	}
	if empty {
		return true
	}
	for nb := 1; nb <= 9; nb++ {
		if isCorrect(*arr, row, col, nb) {
			(*arr)[row][col] = nb
			if solver(arr, len) {
				return true
			} else {
				(*arr)[row][col] = 0
			}
		}
	}
	return false
}

func change(nb rune) int {
	c := 0
	for i := '1'; i <= nb; i++ {
		c++
	}
	return c
}

func slicer(arr []string) [][]int {
	sudoku := make([][]int, 9)
	for r := range sudoku {
		sudoku[r] = make([]int, 9)
	}
	for r, str := range arr {
		for y, ch := range str {
			if ch >= '1' && ch <= '9' {
				sudoku[r][y] = int(ch - '0') // Convert rune to integer
			} else {
				sudoku[r][y] = 0 // Treat non-digit characters as empty cells
			}
		}
	}
	return sudoku
}

func printSudoku(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for s := 0; s < len(arr); s++ {
			fmt.Print(arr[i][s])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func isCorrect(arr [][]int, row int, col int, nb int) bool {
	return !checkR(arr, row, nb) && !checkC(arr, col, nb) && !checkB(arr, row-(row%3), col-(col%3), nb)
}

func checkE(arr []string) bool {
	if len(arr) != 9 {
		return true // Invalid number of rows
	}
	for _, row := range arr {
		if len(row) != 9 {
			return true // Invalid number of columns in a row
		}
	}
	return false // Valid Sudoku grid
}

func checkR(arr [][]int, row int, nb int) bool {
	for col := 0; col < len(arr); col++ {
		if arr[row][col] == nb {
			return true
		}
	}
	return false
}

func checkC(arr [][]int, col int, nb int) bool {
	for row := 0; row < len(arr); row++ {
		if arr[row][col] == nb {
			return true
		}
	}
	return false
}

func checkB(arr [][]int, row int, col int, nb int) bool {
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			if arr[row+r][col+c] == nb {
				return true
			}
		}
	}
	return false
}
