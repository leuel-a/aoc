package main

// PROBLEM: https://adventofcode.com/2025/day/3

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// firstSolution attempts to find the largest number that can be formed by concatenating
// two digits from the input slice `numbers`, where the first digit must appear
// before the second digit in the original slice.
func firstSolution(numbers []int) int {
	firstNumber, secondNumber := -1, -1

	for index, number := range numbers {
		if number > firstNumber && index < len(numbers)-1 {
			firstNumber = number
			secondNumber = -1
		} else if number > secondNumber {
			secondNumber = number
		}
	}

	resultStr := fmt.Sprintf("%v%v", firstNumber, secondNumber)
	result, err := strconv.Atoi(resultStr)
	if err != nil {
		fmt.Printf("Can't convert number %v\n", resultStr)
		return 0
	}

	return result
}

// secondSolution uses Dynamic Programming to find the largest number that can be formed
// by concatenating a sequence of up to `numOfDigits - 1` digits from the input slice,
// followed by the current digit, maximizing the resulting number at each step.
func secondSolution(numbers []int) int {
	numOfDigits := 13 // The maximum number of digits the final concatenated number should have (12 digits + 1 for 0-indexing)
	dp := make([][]int, numOfDigits)

	for digit := range dp {
		dp[digit] = make([]int, len(numbers)+1)
		for index := range dp[digit] {
			dp[digit][index] = -1
		}
	}

	largest := numbers[0]
	for i, value := range numbers {
		largest = max(largest, value)
		dp[1][i] = largest
	}

	for i := 1; i < len(numbers); i++ {
		for digit := 2; digit < numOfDigits; digit++ {
			if dp[digit-1][i-1] == -1 {
				break
			}

			number, _ := strconv.Atoi(fmt.Sprintf("%v%v", dp[digit-1][i-1], numbers[i]))
			dp[digit][i] = max(dp[digit][i-1], number)
		}
	}

	result := -1
	for i := range numbers {
		result = max(dp[numOfDigits-1][i], result)
	}
	return result
}

func RunSolution() {
	inputFile := "input.txt"
	file, err := os.Open(inputFile)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	resultFirst := 0
	resultSecond := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := []rune(scanner.Text())
		numbers := make([]int, len(line))

		for i, rune := range line {
			numbers[i] = int(rune - '0')
		}

		resultFirst += firstSolution(numbers)
		resultSecond += secondSolution(numbers)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}

	fmt.Printf("Result \n\tFirst Part: %v\n\t Second Part: %v\n", resultFirst, resultSecond)
}
