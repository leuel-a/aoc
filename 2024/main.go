package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/leuel-a/aoc/utils"
)

// https://adventofcode.com/2024/day/2
func SolutionPartOne(file *os.File) (int, bool) {
	var safe int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		numbers, err := utils.ConvertToNumArray(parts)

		if err != nil {
			fmt.Printf("Error parsing line `%s`", line)
			os.Exit(1)
		}

		var isValidIncreasing = true
		for i := 1; i < len(numbers); i++ {
			var difference = numbers[i] - numbers[i-1]
			if difference <= 0 || difference > 3 {
				isValidIncreasing = false
				break
			}
		}

		if isValidIncreasing {
			safe++
		}

		var isValidDecreasing = true
		for i := 1; i < len(numbers); i++ {
			var difference = numbers[i] - numbers[i-1]

			if difference > 0 {
				isValidDecreasing = false
				break
			}

			difference = int(math.Abs(float64(difference)))
			if difference <= 0 || difference > 3 {
				isValidDecreasing = false
				break
			}
		}

		if isValidDecreasing {
			safe++
		}
	}

	return safe, true
}

func SolutionPartTwo(os *os.File) {}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening the file: ", err)
		os.Exit(1)
	}
	SolutionPartTwo(file)
}
