package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/leuel-a/aoc/utils"
)

func checkMontonicProperty(numbers []int, omit int, isIncreasing bool) bool {
	for i := range numbers {
		if i == 0 || i == omit {
			continue
		}

		var previous int = numbers[i-1]
		if i-1 == omit {
			if i-2 >= 0 {
				previous = numbers[i-2]
			} else {
				continue
			}
		}

		// TODO: future me this is the one causing the issue if you want to solve it check the isIncreasing boolean
		var difference = numbers[i] - previous
		if !isIncreasing && difference > 0 {
			return false
		}

		difference = int(math.Abs(float64(difference)))
		if difference <= 0 || difference > 3 {
			return false
		}
	}
	return true

}

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

		var isIncreasing = checkMontonicProperty(numbers, -1, true)
		var isDecreasing = checkMontonicProperty(numbers, -1, false)
		if isIncreasing || isDecreasing {
			safe++
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return -1, false
	}

	return safe, true
}

func SolutionPartTwo(file *os.File) (int, bool) {
	var safe int = 0

	var scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		parts := strings.Fields(text)
		nums, err := utils.ConvertToNumArray(parts)

		if err != nil {
			fmt.Printf("Error parsing line %s \n", text)
			return -1, false
		}

		var isIncreasing = checkMontonicProperty(nums, -1, true)
		var isDecreasing = checkMontonicProperty(nums, -1, false)

		if isIncreasing || isDecreasing {
			safe++
		} else {
			for i := range nums {
				isIncreasingOmit := checkMontonicProperty(nums, i, true)
				isDecreasingOmit := checkMontonicProperty(nums, i, false)

				if isIncreasingOmit || isDecreasingOmit {
					safe++
					break
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return -1, false
	}
	return safe, true
}
