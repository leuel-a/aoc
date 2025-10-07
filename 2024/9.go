package main

// https://adventofcode.com/2024/day/9
import (
	"fmt"
	"strconv"
	"strings"

	"github.com/leuel-a/aoc/utils"
)

const FREE_SPACE_VALUE = "."

func buildMemoryRepresentation(content string) ([]string, int) {
	var id int = 0
	var result []string
	numbers, _ := utils.ConvertToNumArray(strings.Split(strings.TrimRight(content, "\r\n"), ""))

	for index, currNumber := range numbers {
		for range currNumber {
			if index%2 == 0 {
				result = append(result, strconv.Itoa(id))
			} else {
				result = append(result, FREE_SPACE_VALUE)
			}
		}

		if index%2 == 0 {
			id++
		}
	}
	return result, id - 1
}

func getLastUsedSpace(memory []string) int {
	var lastUsedSpace = -1
	for i := len(memory) - 1; i >= 0; i-- {
		if memory[i] != FREE_SPACE_VALUE {
			lastUsedSpace = i
			break
		}
	}
	return lastUsedSpace
}

func getFirstFreeSpace(memory []string) int {
	var firstFreeSpace = -1
	for i := range memory {
		if memory[i] == FREE_SPACE_VALUE {
			firstFreeSpace = i
			break
		}
	}
	return firstFreeSpace
}

func getLastUsedRange(memory []string, id int) (int, int) {
	var startIndex int = -1
	var lastIndex int = -1
	for i, value := range memory {
		if value == strconv.Itoa(id) && startIndex == -1 {
			startIndex = i
		}

		if startIndex != -1 && value == strconv.Itoa(id) {
			lastIndex = i
		}
	}
	return startIndex, lastIndex
}

func rangeIsFreeSpace(memory []string, start int, end int) bool {
	for i := start; i <= end; i++ {
		if memory[i] != FREE_SPACE_VALUE {
			return false
		}
	}
	return true
}

func findFreeSpace(memory []string, length int) (int, int, bool) {
	for i := range memory {
		if i-length >= 0 && rangeIsFreeSpace(memory, i-length, i) {
			return i - length, i, true
		}
	}
	return -1, -1, false
}

func calculateCheckSum(memory []string) int {
	var checksum = 0
	for index, id := range memory {
		if id == FREE_SPACE_VALUE {
			continue
		}
		currNumber, _ := strconv.Atoi(id)
		checksum = checksum + (index * currNumber)
	}
	return checksum
}

func DayNineSolutionPartTwo(content string) {
	memory, id := buildMemoryRepresentation(content)

	// INFO: this is slow as FUCCKKK
	for id > 0 {
		usedStart, usedEnd := getLastUsedRange(memory, id)
		length := usedEnd - usedStart

		freeStart, freeEnd, isFreeSpace := findFreeSpace(memory, length)
		// fmt.Println("Used (", usedStart, ", ", usedEnd, ")", "Free (", freeStart, ", ", freeEnd, ")", "ID", id)

		if isFreeSpace && usedStart > freeStart {
			for i := freeStart; i <= freeEnd; i++ {
				memory[i] = memory[usedStart]
			}
			for i := usedStart; i <= usedEnd; i++ {
				memory[i] = FREE_SPACE_VALUE
			}
		}
		id--
		// fmt.Println(memory)
		// fmt.Println()
	}
	// fmt.Println(memory)
	fmt.Println(calculateCheckSum(memory))
}

func DayNineSolutionPartOne(content string) {
	memory, _ := buildMemoryRepresentation(content)
	var lastUsedSpace int = getLastUsedSpace(memory)
	var firstFreeSpace int = getFirstFreeSpace(memory)

	for firstFreeSpace < lastUsedSpace {
		memory[firstFreeSpace] = memory[lastUsedSpace]
		memory[lastUsedSpace] = FREE_SPACE_VALUE

		lastUsedSpace = getLastUsedSpace(memory)
		firstFreeSpace = getFirstFreeSpace(memory)
	}

	fmt.Println(calculateCheckSum(memory))
}
