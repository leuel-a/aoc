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
