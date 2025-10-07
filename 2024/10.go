package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/leuel-a/aoc/utils"
)

var visited [][2]int

func isInBound(grid [][]int, row int, column int) bool {
	return (row >= 0 && row < len(grid)) && (column >= 0 && column < len(grid[0]))
}

func valueNotInVisited(row int, column int) bool {
	for _, value := range visited {
		if row == value[0] && column == value[1] {
			return false
		}
	}
	return true
}

func countDistnictTrails(grid [][]int, row int, column int) int {
	if grid[row][column] == 9 {
		return 1
	}

	var count = 0
	var directions = [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for _, direction := range directions {
		nextRow := row + direction[0]
		nextColumn := column + direction[1]

		if isInBound(grid, nextRow, nextColumn) && grid[nextRow][nextColumn]-grid[row][column] == 1 {
			// fmt.Printf("(%d) ->> (%d) (%d, %d)\n", grid[row][column], grid[nextRow][nextColumn], nextRow, nextColumn)
			count += countDistnictTrails(grid, nextRow, nextColumn)
		}
	}

	return count
}

func countTrails(grid [][]int, row int, column int) int {
	if grid[row][column] == 9 && valueNotInVisited(row, column) {
		visited = append(visited, [2]int{row, column})
		return 1
	}

	var count = 0
	var directions = [4][2]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	for _, direction := range directions {
		nextRow := row + direction[0]
		nextColumn := column + direction[1]

		if isInBound(grid, nextRow, nextColumn) && grid[nextRow][nextColumn]-grid[row][column] == 1 {
			// fmt.Printf("(%d) ->> (%d) (%d, %d)\n", grid[row][column], grid[nextRow][nextColumn], nextRow, nextColumn)
			count += countTrails(grid, nextRow, nextColumn)
		}
	}

	return count
}

func DayTenSolutionPartOne(file *os.File) {
	var grid [][]int
	var scanner = bufio.NewScanner(file)
	var startPoints [][2]int

	for scanner.Scan() {
		row, _ := utils.ConvertToNumArray(strings.Split(strings.TrimRight(scanner.Text(), "\r\n"), ""))
		grid = append(grid, row)
	}

	for i, row := range grid {
		for j, value := range row {
			if value == 0 {
				startPoints = append(startPoints, [2]int{i, j})
			}
		}
	}

	var trails = 0
	for _, point := range startPoints {
		visited = nil
		trails += countTrails(grid, point[0], point[1])
		// fmt.Printf("START (%d, %d) TRAILS %d\n", point[0], point[1], trails)
	}

	if err := scanner.Err(); err != nil {
		panic("(error) the was an error when parsing an error")
	}

	fmt.Println(trails)
}

func DayTenSolutionPartTwo(file *os.File) {
	var grid [][]int
	var scanner = bufio.NewScanner(file)
	var startPoints [][2]int

	for scanner.Scan() {
		row, _ := utils.ConvertToNumArray(strings.Split(strings.TrimRight(scanner.Text(), "\r\n"), ""))
		grid = append(grid, row)
	}

	for i, row := range grid {
		for j, value := range row {
			if value == 0 {
				startPoints = append(startPoints, [2]int{i, j})
			}
		}
	}

	var trails = 0
	for _, point := range startPoints {
		trails += countDistnictTrails(grid, point[0], point[1])
	}

	if err := scanner.Err(); err != nil {
		panic("(error) the was an error when parsing an error")
	}

	fmt.Println(trails)
}
