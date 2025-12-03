package main

// PROBLEM: https://adventofcode.com/2025/day/2

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CheckIsInvalidIDPart1 implements the logic intended for the "First Part Solution,"
// which checks for a simple, full-half repetition (e.g., 123123).
// It returns the ID as an integer if it matches the criteria, otherwise -1.
func CheckIsInvalidIDPart1(id string) int {
	if len(id)%2 != 0 {
		return -1
	}

	halfLen := len(id) / 2
	firstPart := id[:halfLen]
	secondPart := id[halfLen:]

	if firstPart == secondPart {
		fmt.Printf("found Part 1 invalid id `%s` (simple repetition)\n", id)
		convertedId, _ := strconv.Atoi(id)
		return convertedId
	}

	return -1
}

// CheckIsInvalidIDPart2 implements the logic intended for the "Second Part Solution,"
// which checks for any repeating substring pattern that divides the full ID (e.g., 101010, 5555).
// This was the active logic in the original provided function.
// It returns the ID as an integer if it matches the criteria, otherwise -1.
func CheckIsInvalidIDPart2(id string) int {
	for i := len(id) / 2; i > 0; i-- {
		fmt.Printf("\tcurrent interval for id: `%s`, is %d\n", id, i)

		var value = ""
		var isInvalid, seenFullIndex = true, false

		for j := 0; j < len(id) && j+i <= len(id); j += i {
			currentSub := id[j : j+i]

			if value == "" {
				value = currentSub
			} else {
				if value != currentSub {
					isInvalid = false
					break
				}
			}

			if j+i == len(id) {
				seenFullIndex = true
			}
		}

		if isInvalid && seenFullIndex {
			fmt.Printf("found Part 2 invalid id `%s` (pattern length %d)\n", id, i)
			convertedId, _ := strconv.Atoi(id)
			return convertedId
		}
	}
	return -1
}

func Solution() {
	var inputFile = "./input.txt"
	file, err := os.ReadFile(inputFile)
	if err != nil {
		fmt.Printf("Error reading file %s: %v\n", inputFile, err)
	}

	content := strings.ReplaceAll(string(file), "\n", "")
	ids := strings.Split(content, ",")

	var result = 0
	for _, idRange := range ids {
		var start, end = -1, -1

		rangeParts := strings.Split(idRange, "-")
		if len(rangeParts) == 2 {
			start, _ = strconv.Atoi(rangeParts[0])
			end, _ = strconv.Atoi(rangeParts[1])
		}

		for i := start; i <= end; i++ {
			currentID := strconv.Itoa(i)

			// The original code used the logic from Part 2, so we call that here.
			// if invalid := checkIsInvalidIDPart2(currentID); invalid != -1 {
			// 	result += invalid
			// }

			// If the "First Part Solution" (simple repetition) was intended, it would look like this:
			if invalid := CheckIsInvalidIDPart1(currentID); invalid != -1 {
				result += invalid
			}
		}
	}
	fmt.Printf("Result: %d\n", result)
}
