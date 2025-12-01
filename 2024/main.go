package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func DayThirteenSolutionPartOne(file *os.File) {
	var buttonA, buttonB, target *Point
	var scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		var text = scanner.Text()
		if text == "" {
			fmt.Printf("Button A: %d, %d\n", buttonA.X, buttonA.Y)
			fmt.Printf("Button B: %d, %d\n", buttonB.X, buttonB.Y)
			fmt.Printf("Target: %d, %d\n", target.X, target.Y)
			fmt.Println()

			buttonA, buttonB, target = nil, nil, nil
		} else {
			var expression = regexp.MustCompile(`(?:Button |Prize)\w?: X[\+=](\d+), Y[\+=](\d+)`)
			var matches = expression.FindStringSubmatch(text)

			var firstNumber, _ = strconv.Atoi(matches[1])
			var secondNumber, _ = strconv.Atoi(matches[2])

			if buttonA == nil {
				buttonA = &Point{X: firstNumber, Y: secondNumber}
			} else if buttonB == nil {
				buttonB = &Point{X: firstNumber, Y: secondNumber}
			} else {
				target = &Point{X: firstNumber, Y: secondNumber}
			}
		}
	}
}

func main() {
	var INPUT_FILE = "input.txt"

	file, err := os.Open(INPUT_FILE)
	if err != nil {
		panic("Cannot Open Test File")
	}

	DayThirteenSolutionPartOne(file)
}
