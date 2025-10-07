package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/leuel-a/aoc/utils"
)

func generateCombinations(count int) []string {
	var result []string
	var generate func(prefix string, index int)

	generate = func(prefix string, index int) {
		if index == count {
			result = append(result, prefix)
			return
		}

		generate(prefix+"+", index+1) // pick +
		generate(prefix+"*", index+1) // pick *

		// NOTE: remove this for the part one solution of the problem
		generate(prefix+"|", index+1) // pick |
	}

	generate("", 0)
	return result
}

func applyOperation(operation string, numbers []int) int {
	var result int = numbers[0]
	for i, char := range operation {
		if string(char) == "+" {
			result = result + numbers[i+1]
		} else if string(char) == "*" {
			result = result * numbers[i+1]
		} else {
			// NOTE: this needs to be removed for the part one solution of the problem
			concatenated := strconv.Itoa(result) + strconv.Itoa(numbers[i+1])
			convertedNumber, _ := strconv.Atoi(concatenated)
			result = convertedNumber
		}
	}
	return result
}

func DaySevenSolution(file *os.File) {
	totalCalibrationResult := 0
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")

		if len(parts) != 2 {
			fmt.Println("(error) the current line is not value", scanner.Text())
			os.Exit(1)
		}

		resultNumber, _ := strconv.Atoi(parts[0])
		numbers, _ := utils.ConvertToNumArray(strings.Split(strings.TrimSpace(parts[1]), " "))
		operations := generateCombinations(len(numbers) - 1)

		for _, operation := range operations {
			operationResult := applyOperation(operation, numbers)

			if operationResult == resultNumber {
				totalCalibrationResult += resultNumber
				break
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("(error) unable to parse the file, scanner failed")
		os.Exit(1)
	}
	fmt.Println(totalCalibrationResult)
}
