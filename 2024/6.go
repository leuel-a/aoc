package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/leuel-a/aoc/utils"
)

type GridPoint struct {
	row    int
	column int
}

// INFO:Global Variables in Golang are package-level
// https://go101.org/article/blocks-and-scopes.html
// https://go.dev/ref/spec
var guardsNotations = []string{"^", ">", "<", "v"}

var directions = map[string]string{
	"^": ">",
	">": "v",
	"v": "<",
	"<": "^",
}

var movement = map[string]GridPoint{
	"^": {-1, 0},
	">": {0, 1},
	"v": {1, 0},
	"<": {0, -1},
}

func readGridFromFile(file *os.File) ([][]string, bool) {
	var grid [][]string
	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		grid = append(grid, strings.Split(scanner.Text(), ""))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("(error) something went wrong while trying to read the file", err)
		os.Exit(1)
	}
	return grid, true
}

func findGuardInGrid(grid [][]string) GridPoint {
	var guardLocation GridPoint
	for rowIndex, currRow := range grid {
		for columnIndex, currValue := range currRow {
			if slices.Contains(guardsNotations, currValue) {
				guardLocation = GridPoint{rowIndex, columnIndex}
			}
		}
	}
	return guardLocation
}

func gridPointInRange(gridPoint GridPoint, row int, column int) bool {
	return (gridPoint.row >= 0 && gridPoint.row < row) && (gridPoint.column >= 0 && gridPoint.column < column)
}

func markGuardPath(grid [][]string, guardLocation GridPoint) ([][]string, int) {
	var newGrid = utils.CloneGrid(grid)
	var leftGrid bool = false

	var n = len(newGrid)
	var m = len(newGrid[0])
	var currPoint = guardLocation
	var guardDirection = newGrid[currPoint.row][currPoint.column]

	for leftGrid != true {
		newGrid[currPoint.row][currPoint.column] = "X"

		nextMove := movement[guardDirection]
		nextGridPoint := GridPoint{currPoint.row + nextMove.row, currPoint.column + nextMove.column}
		if !gridPointInRange(nextGridPoint, n, m) {
			leftGrid = true
		} else {
			nextGridPointValue := newGrid[nextGridPoint.row][nextGridPoint.column]

			if nextGridPointValue == "#" {
				guardDirection = directions[guardDirection]
				currPoint = GridPoint{currPoint.row + movement[guardDirection].row, currPoint.column + movement[guardDirection].column}
			} else {
				currPoint = nextGridPoint
			}
		}
	}

	var result int = 0
	for _, row := range newGrid {
		for _, value := range row {
			if value == "X" {
				result++
			}
		}
	}
	return newGrid, result
}

// ^ v -->> |
// > < -->> -
func markPathForObstruction(direction string) string {
	if direction == "^" || direction == "v" {
		return "|"
	}
	return "-"
}

func createObstructionsOnGrid(grid [][]string, guardLocation GridPoint) int {
	var result = 0
	var newGrid = utils.CloneGrid(grid)
	var leftGrid bool = false

	var n = len(newGrid)
	var m = len(newGrid[0])
	var currPoint = guardLocation
	var visitedStates = []string{"|", "-", "+", "^"}
	var guardDirection = newGrid[currPoint.row][currPoint.column]
	var directionChanges []string

	for leftGrid != true {
		if currPoint != guardLocation {
			newGrid[currPoint.row][currPoint.column] = markPathForObstruction(guardDirection)
		}

		nextMove := movement[guardDirection]
		nextGridPoint := GridPoint{currPoint.row + nextMove.row, currPoint.column + nextMove.column}
		if !gridPointInRange(nextGridPoint, n, m) {
			leftGrid = true
		} else {
			nextGridPointValue := newGrid[nextGridPoint.row][nextGridPoint.column]

			if nextGridPointValue == "#" {
				guardDirection = directions[guardDirection]
				directionChanges = append(directionChanges, guardDirection)
				newGrid[currPoint.row][currPoint.column] = "+"
				currPoint = GridPoint{currPoint.row + movement[guardDirection].row, currPoint.column + movement[guardDirection].column}

			} else {
				if slices.Contains(visitedStates, nextGridPointValue) {
					possibleObstructionPoint := GridPoint{nextGridPoint.row + nextMove.row, nextGridPoint.column + nextMove.column}
					if gridPointInRange(possibleObstructionPoint, n, m) {
						possibleObstructionValue := newGrid[possibleObstructionPoint.row][possibleObstructionPoint.column]
						if possibleObstructionValue != "#" && len(directionChanges)%2 == 1 {
							result += 1
						}
					}
				}
				currPoint = nextGridPoint
			}
		}
	}
	return result
}

// https://adventofcode.com/2024/day/6
func DaySixSolutionPartOne(file *os.File) (int, [][]string) {
	grid, _ := readGridFromFile(file)

	guardLocation := findGuardInGrid(grid)
	markedGrid, result := markGuardPath(grid, guardLocation)
	return result, markedGrid
}

func DaySixSolutionPartTwo(file *os.File) int {
	grid, _ := readGridFromFile(file)
	guardLocation := findGuardInGrid(grid)
	return createObstructionsOnGrid(grid, guardLocation)
}
