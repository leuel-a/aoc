package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

// https://adventofcode.com/2024/day/1
func DayOneSolutionPartOne(file *os.File) (int, bool) {
	scanner := bufio.NewScanner(file)
	var leftNumbers, rightNumbers []int

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		for i, s := range parts {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Conversion Error: ", err)
				return -1, false
			}

			if i%2 == 0 {
				leftNumbers = append(leftNumbers, n)
			} else {
				rightNumbers = append(rightNumbers, n)
			}
		}
	}

	sort.Ints(leftNumbers)
	sort.Ints(rightNumbers)

	size := len(leftNumbers)

	var result float64
	for i := range size {
		result += math.Abs(float64(leftNumbers[i] - rightNumbers[i]))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
		return -1, false
	}
	return int(result), true
}

func DayOneSolutionPartTwo(file *os.File) (int, bool) {
	scanner := bufio.NewScanner(file)
	var leftNumbers []int
	var rightNumbersCount = make(map[int]int)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		for i, s := range parts {
			n, err := strconv.Atoi(s)
			if err != nil {
				fmt.Println("Coversion Error: ", err)
				return -1, false
			}

			if i%2 == 0 {
				leftNumbers = append(leftNumbers, n)
			} else {
				val, ok := rightNumbersCount[n]
				if !ok {
					rightNumbersCount[n] = 1
				} else {
					rightNumbersCount[n] = val + 1
				}
			}
		}
	}

	var result int
	for _, num := range leftNumbers {
		count, ok := rightNumbersCount[num]
		if ok {
			result += (num * count)
		}
	}

	return result, true
}
