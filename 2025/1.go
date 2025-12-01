package main

// PROBLEM: https://adventofcode.com/2025/day/1

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var inputFile = "input.txt"

// DayOneSolutionPartOne solves the first part of the puzzle.
// It counts the number of times the dial is left pointing exactly at 0
// *after* any rotation in the sequence is complete.
func DayOneSolutionPartOne() {
	// Open the puzzle input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file for Part One: %v\n", err)
		return
	}
	defer file.Close()

	var current = 50 // The dial starts by pointing at 50.
	var count = 0    // Counter for the number of times the final position is 0.
	var scanner = bufio.NewScanner(file)
	var expression = regexp.MustCompile(`([LR])(\d+)`)

	// Process each rotation command
	for scanner.Scan() {
		var text = scanner.Text()
		var matches = expression.FindStringSubmatch(text)

		if len(matches) < 3 {
			continue // Skip malformed lines
		}

		// distance is how many clicks to rotate
		difference, _ := strconv.Atoi(matches[2])

		// Apply the rotation
		if matches[1] == "L" {
			// Left rotation (toward lower numbers).
			// Formula: (current - difference % 100 + 100) % 100 ensures a non-negative result.
			current = (current - difference%100 + 100) % 100
		} else {
			// Right rotation (toward higher numbers).
			// Formula: (current + difference) % 100
			current = (current + difference) % 100
		}

		// Check if the final position after the rotation is 0
		if current == 0 {
			count += 1
		}
	}

	fmt.Printf("--- Day 1: Secret Entrance ---\n")
	fmt.Printf("Part One Answer (Final Position Count): %d\n", count)
}

// DayOneSolutionPartTwo solves the second part of the puzzle.
// It counts the number of times *any click* causes the dial to point at 0,
// regardless of whether it happens during a rotation or at the end of one.
func DayOneSolutionPartTwo() {
	// Open the puzzle input file
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file for Part Two: %v\n", err)
		return
	}
	defer file.Close()

	var current = 50      // The dial starts by pointing at 50.
	var totalZeroHits = 0 // Counter for *all* times 0 is hit (during or at the end).
	var scanner = bufio.NewScanner(file)
	var expression = regexp.MustCompile(`([LR])(\d+)`)

	// Process each rotation command
	for scanner.Scan() {
		var text = scanner.Text()
		var matches = expression.FindStringSubmatch(text)

		if len(matches) < 3 {
			continue // Skip malformed lines
		}

		difference, _ := strconv.Atoi(matches[2])
		direction := matches[1]

		if direction == "L" {
			// **Calculate Zero Hits for Left Rotation**
			// Clicks needed to hit the first 0 (e.g., from 50 to 0 is 50 clicks).
			// The boundary is 0 -> 99. The first 0 is hit at the ClicksToFirstZero step.
			clicksToFirstZero := current
			if current == 0 {
				clicksToFirstZero = 100 // If starting at 0, the next 0 is 100 clicks away.
			}

			if difference >= clicksToFirstZero {
				// The first 0 is hit.
				totalZeroHits++
				remainingClicks := difference - clicksToFirstZero
				// Every subsequent 100 clicks also hits 0.
				totalZeroHits += remainingClicks / 100
			}

			// Update the current position
			current = (current - difference%100 + 100) % 100

		} else { // direction == "R"
			// **Calculate Zero Hits for Right Rotation**
			// Clicks needed to hit the first 0 (e.g., from 50 past 99 to 0 is 50 clicks).
			// The boundary is 99 -> 0. The first 0 is hit at the ClicksToFirstZero step.
			clicksToFirstZero := 100 - current
			if current == 0 {
				clicksToFirstZero = 100 // If starting at 0, the next 0 is 100 clicks away.
			}

			if difference >= clicksToFirstZero {
				// The first 0 is hit.
				totalZeroHits++
				remainingClicks := difference - clicksToFirstZero
				// Every subsequent 100 clicks also hits 0.
				totalZeroHits += remainingClicks / 100
			}

			// Update the current position
			current = (current + difference) % 100
		}
	}

	fmt.Printf("Part Two Answer (All Click Zero Count): %d\n", totalZeroHits)
}

// Main function to run both solutions
func run() {
	// Call the solutions in order
	DayOneSolutionPartOne()
	fmt.Println("--------------------------------")
	DayOneSolutionPartTwo()
}
