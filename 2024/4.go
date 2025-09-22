package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func buildStringCol(grid [][]rune, column int, startRow int) string {
	var resultString []rune
	for i := startRow; i <= startRow+3; i++ {
		resultString = append(resultString, grid[i][column])
	}
	return string(resultString)
}

func buildStringDiag(grid [][]rune, row int, column int, upDirection bool) string {
	var result []rune
	for k := range 4 {
		if upDirection {
			result = append(result, grid[row-k][column+k])
		} else {
			result = append(result, grid[row+k][column+k])
		}
	}
	return string(result)
}

func DayFourSolutionPartOne(file *os.File) (int, bool) {
	var scanner = bufio.NewScanner(file)

	var grid [][]rune
	for scanner.Scan() {
		var text = scanner.Text()
		var chars = []rune(text)
		grid = append(grid, chars)

	}

	var count = 0
	var n = len(grid)
	var m = len(grid[0])

	for i := range n {
		for j := range m {
			var currString string = ""

			// horizontal
			if j+3 < m {
				currString = string(grid[i][j : j+4])
				if currString == "XMAS" || currString == "SAMX" {
					count++
				}
			}

			// vertical
			if i+3 < n {
				currString = buildStringCol(grid, j, i)
				if currString == "XMAS" || currString == "SAMX" {
					count++
				}
			}

			// diagnonal
			if i+3 < n && j+3 < m {
				currString = buildStringDiag(grid, i, j, false)
				if currString == "XMAS" || currString == "SAMX" {
					count++
				}
			}

			if i-3 >= 0 && j+3 < m {
				currString = buildStringDiag(grid, i, j, true)
				if currString == "XMAS" || currString == "SAMX" {
					count++
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("(ERROR) unable to scan the file")
	}
	return count, true
}

func buildXMASString(grid [][]rune, row int, column int) string {
	var offsets = [][2]int{
		{0, 0}, // grid[row][column]
		{2, 0}, // grid[row+2][column]
		{1, 1}, // grid[row+1][column+1]
		{0, 2}, // grid[row][column+2]
		{2, 2}, // grid[row+2][column+2]
	}
	var result []rune

	result = make([]rune, 0, len(offsets))
	for _, offset := range offsets {
		currRow, currCol := row+offset[0], column+offset[1]
		result = append(result, grid[currRow][currCol])
	}
	return string(result)
}

func DayFourSolutionPartTwo(file *os.File) (int, bool) {
	var scanner = bufio.NewScanner(file)

	var grid [][]rune
	for scanner.Scan() {
		var text = scanner.Text()
		var chars = []rune(text)
		grid = append(grid, chars)
	}

	var count = 0
	var n = len(grid)
	var m = len(grid[0])

	for i := range n {
		for j := range m {
			if i+2 < n && j+2 < n {
				candidates := []string{"MMASS", "SMASM", "SSAMM", "MSAMS"}
				if slices.Contains(candidates, buildXMASString(grid, i, j)) {
					count++
				}
			}
		}
	}

	return count, true
}
