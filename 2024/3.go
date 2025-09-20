package main

import (
	"strconv"
	"unicode"

	"github.com/leuel-a/aoc/utils"
)

func evaluateStack(stack utils.Stack) int {
	stack.Pop()
	firstNumberString, _ := stack.Pop()
	stack.Pop()
	secondNumberString, _ := stack.Pop()
	stack.Pop()

	firstNumber, _ := strconv.Atoi(firstNumberString)
	secondNumber, _ := strconv.Atoi(secondNumberString)

	return firstNumber * secondNumber
}

func getCurrentValue(index int, content string) (int, bool) {
	var result int = 0
	var stack utils.Stack
	if string(content[index]) != "(" {
		return -1, false
	}

	stack.Push("(")
	for i := index + 1; i < len(content); i++ {
		if string(content[i]) == ")" {
			stack.Push(")")
			break
		} else if string(content[i]) == "," {
			if len(stack) != 2 {
				return -1, false
			}
			stack.Push(string(content[i]))
		} else if unicode.IsNumber(rune(content[i])) {
			if top, _ := stack.Peek(); len(stack) == 1 && top == "(" {
				stack.Push(string(content[i]))
			} else if top, _ := stack.Peek(); len(stack) == 3 && top == "," {
				stack.Push(string(content[i]))
			} else {
				top, _ := stack.Pop()
				stack.Push(top + string(content[i]))
			}
		} else {
			return -1, false
		}
	}
	result = evaluateStack(stack)
	return result, true
}

func DayThreeSolutionPartOne(content string) (int, bool) {
	var result = 0

	for i := 3; i < len(content); i++ {
		curr_candidate := content[i-3 : i]
		if curr_candidate == "mul" {
			value, ok := getCurrentValue(i, content)
			if ok {
				result += value
			}
		}
	}
	return result, true
}

func DayThreeSolutionPartTwo(content string) (int, bool) {
	var result = 0
	var evaluate = true

	for i := 3; i < len(content); i++ {
		var curr_candidate string = ""

		if i > 6 {
			curr_candidate = content[i-7 : i]
			if curr_candidate == "don't()" {
				evaluate = false
				continue
			}
		}

		if i > 3 {
			curr_candidate = content[i-4 : i]
			if curr_candidate == "do()" {
				evaluate = true
				continue
			}
		}

		if evaluate {
			curr_candidate = content[i-3 : i]
			if curr_candidate == "mul" {
				value, ok := getCurrentValue(i, content)
				if ok {
					result += value
				}
			}
		}

	}
	return result, true
}
