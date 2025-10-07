package utils

import "strconv"

// ConvertToNumArray converts a slice of string representations of numbers
// into a slice of integers. If any string in the input slice cannot be
// converted to an integer, the function returns nil.
func ConvertToNumArray(input []string) ([]int, error) {
	var result = make([]int, len(input))

	for i, s := range input {
		num, err := strconv.Atoi(s)

		if err != nil {
			return nil, err
		}
		result[i] = num
	}
	return result, nil
}

// CloneGrid creates a deep copy of a two-dimensional slice of integers (a grid).
func CloneGrid(originalGrid [][]string) [][]string {
	newGrid := make([][]string, len(originalGrid))

	for i, row := range originalGrid {
		newGrid[i] = make([]string, len(row))
		copy(newGrid[i], row)
	}
	return newGrid
}
