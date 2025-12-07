package main

import (
	"bufio"
	"fmt"
	"os"
)

type Direction struct {
	X int
	Y int
}

// inBound checks if a given (row, col) coordinate pair is within the valid boundaries of the provided 2D grid.
func inBound(grid [][]rune, row int, col int) bool {
	return (row >= 0 && row < len(grid)) && (col >= 0 && col < len(grid[0]))
}

// getNumOfAccessedPapers iterates through the grid to identify and modify specific elements.
// It checks each cell containing '@' and counts its '@' neighbors (including diagonals).
// If an '@' cell has less than 4 '@' neighbors, it is changed to '.' and contributes 1 to the result count.
// This function modifies the input grid in place.
func getNumOfAccessedPapers(grid [][]rune) int {
	directions := []Direction{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: 1, Y: 0},
		{X: -1, Y: 0},
		{X: 1, Y: 1},
		{X: 1, Y: -1},
		{X: -1, Y: -1},
		{X: -1, Y: 1},
	}

	result := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == '@' {
				count := 0
				for _, direction := range directions {
					if inBound(grid, i+direction.X, j+direction.Y) && grid[i+direction.X][j+direction.Y] == '@' {
						count++
					}
				}

				if count < 4 {
					grid[i][j] = '.'
					result += 1
				}
			}
		}

	}
	return result
}

// firstSolutionDayFour serves as the entry point for the "first part" of the problem (likely an Advent of Code challenge).
// It performs a single iteration of the paper access/modification logic on the grid.
func firstSolutionDayFour(grid [][]rune) int {
	return getNumOfAccessedPapers(grid)
}

// secondSolutionDayFour serves as the entry point for the "second part" of the problem.
// It repeatedly applies the getNumOfAccessedPapers logic until no more '@' cells are modified in a single iteration (i.e., currCount == 0).
// It sums up the total number of modified cells across all iterations. This pattern suggests a simulation or steady-state problem.
func secondSolutionDayFour(grid [][]rune) int {
	result := 0
	for {
		currCount := getNumOfAccessedPapers(grid)
		if currCount == 0 {
			break
		}
		result += currCount
	}
	return result
}

func runSolution() {
	inputFile := "input.txt"
	file, _ := os.Open(inputFile)
	scanner := bufio.NewScanner(file)

	var grid [][]rune
	for scanner.Scan() {
		text := scanner.Text()
		grid = append(grid, []rune(text))
	}

	fmt.Printf("Result \n\tFirst Part: %v \n\tSecond Part: %v\n", firstSolutionDayFour(grid), secondSolutionDayFour(grid))
}
