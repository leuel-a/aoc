package main

// https://adventofcode.com/2024/day/5 -- Print Queue

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/leuel-a/aoc/utils"
)

func isValidSequence(sequence []int, suffixList map[int][]int) bool {
	for i, value := range sequence {
		if i > 0 {
			for j := i - 1; j >= 0; j-- {
				if slices.Contains(suffixList[value], sequence[j]) {
					return false
				}
			}
		}
	}
	return true
}

func fixCurrentSequence(sequence []int, suffixList map[int][]int) []int {
	validSequence := append([]int(nil), sequence...)

	for i := 1; i < len(validSequence); i++ {
		var currIndex = i
		var currValue = validSequence[i]
		for j := i - 1; j >= 0; j-- {
			var candidateValue = validSequence[j]
			if slices.Contains(suffixList[currValue], candidateValue) {
				// we need to swap the two indices since they are not in the valid order
				validSequence[currIndex] = candidateValue
				validSequence[j] = currValue
				currIndex = j
			}
		}
	}

	return validSequence
}

func DayFiveSolutionPartOne(file *os.File) (int, bool) {
	var result = 0
	var suffixList = make(map[int][]int)
	var sequences [][]int
	var buildGraph = true
	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		var text = scanner.Text()

		if text == "" {
			buildGraph = false
			continue
		}

		if buildGraph {
			values, _ := utils.ConvertToNumArray(strings.Split(text, "|"))
			numbers := [2]int{values[0], values[1]}
			suffixList[numbers[0]] = append(suffixList[numbers[0]], numbers[1])
		} else {
			currSequence, _ := utils.ConvertToNumArray(strings.Split(text, ","))
			sequences = append(sequences, currSequence)
		}
	}

	for _, sequence := range sequences {
		if isValidSequence(sequence, suffixList) {
			result += sequence[len(sequence)/2]
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("(ERROR) scanner failed", err)
		return -1, false
	}

	return result, true
}

func DayFiveSolutionPartTwo(file *os.File) (int, bool) {
	var result int = 0
	var scanner = bufio.NewScanner(file)
	var sequences [][]int
	var suffixList = make(map[int][]int)
	var buildGraph = true

	for scanner.Scan() {
		var text = scanner.Text()

		if text == "" {
			buildGraph = false
			continue
		}

		if buildGraph {
			values, _ := utils.ConvertToNumArray(strings.Split(scanner.Text(), "|"))
			numbers := [2]int{values[0], values[1]}
			suffixList[numbers[0]] = append(suffixList[numbers[0]], numbers[1])
		} else {
			currSequence, _ := utils.ConvertToNumArray(strings.Split(scanner.Text(), ","))
			sequences = append(sequences, currSequence)
		}
	}

	for _, sequence := range sequences {
		if isValidSequence(sequence, suffixList) {
			continue
		}

		validSequence := fixCurrentSequence(sequence, suffixList)
		result += validSequence[len(validSequence)/2]
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("(ERROR) scanner failed", err)
		return -1, false
	}
	return result, true
}
