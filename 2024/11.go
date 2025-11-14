package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// -- If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
// -- If the stone is engraved with a number that has an even number of digits, it is replaced by two stones.
//
//	The left half of the digits are engraved on the new left stone, and the right half of the digits are engraved on the new right stone.
//	(The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)
//
// -- If none of the other rules apply, the stone is replaced by a new stone; the old stone's number multiplied by 2024 is engraved on the new stone.

func removeLeadingZeros(candidate string) string {
	if string(candidate[0]) != "0" {
		return candidate
	}

	var firstNonZeroIndex = -1
	// find the first non-zero character
	for index, value := range candidate {
		if string(value) != "0" {
			firstNonZeroIndex = index
			break
		}
	}

	// slice the string from that character till the end
	if firstNonZeroIndex == -1 {
		return "0"
	}

	nonLeadingZeroNumber := candidate[firstNonZeroIndex:]
	if len(nonLeadingZeroNumber) == 0 {
		return "0"
	}
	return nonLeadingZeroNumber
}

func digitCountOfProduct(a, b float64) int {
	if a == 0 || b == 0 {
		return 1
	}
	digits := math.Floor(math.Log10(a)+math.Log10(b)) + 1
	return int(digits)
}

func generateNextLengths(lengths []int) []int {
	var nextLengths []int

	for _, n := range lengths {
		if n == 1 {
			nextLengths = append(nextLengths, 4)
		} else if n%2 == 0 {
			nextLengths = append(nextLengths, n/2, n/2)
		} else {
			nextLengths = append(nextLengths, n*4)
		}
	}
	return nextLengths
}

func DayElevenSolutionPartTwo(content string) {
	var numbers = strings.Split(strings.TrimRight(content, "\r\n"), " ")
	var lengths = make([]int, len(numbers))

	for i, n := range numbers {
		lengths[i] = len(n)
	}

	fmt.Println(lengths)
	for i := 1; i <= 6; i++ {
		lengths = generateNextLengths(lengths)
		fmt.Println(lengths)
	}
	fmt.Println(len(lengths))
}

func DayElevenSolutionPartOne(content string) {
	var numbers = strings.Split(strings.TrimRight(content, "\r\n"), " ")

	var currNumbers []string
	for i := 1; i <= 75; i++ {
		currNumbers = nil
		for _, value := range numbers {
			if value == "0" {
				currNumbers = append(currNumbers, "1")
			} else if len(value)%2 == 0 {
				currNumbers = append(currNumbers, removeLeadingZeros(value[:len(value)/2]), removeLeadingZeros(value[len(value)/2:]))
			} else {
				numericValue, _ := strconv.Atoi(value)
				currNumbers = append(currNumbers, strconv.Itoa(numericValue*2024))
			}
		}

		numbers = currNumbers
	}
	fmt.Println(len(numbers))
}
