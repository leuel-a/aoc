#!/usr/bin/env python3
"""
Advent Of Code[2022] - Day 1

Problem Link: https://adventofcode.com/2022/day/1
Difficulty: Easy
"""

INPUT_FILE = 'input.txt'
SEPARATOR = '\n'

def solutionPartOne():
    with open(INPUT_FILE, "r") as file:
        content = file.read()
        lines = content.split(SEPARATOR)

        max_calories = 0
        curr_calories = 0

        for line in lines:
            if line.strip(SEPARATOR) == "":
                curr_calories = 0
            else:
                curr_calories += int(line)
                max_calories = max(max_calories, curr_calories)
        return max_calories


def solutionPartTwo():
    with open(INPUT_FILE, "r") as file:
        content = file.read()
        lines = content.split(SEPARATOR)

        calories = []
        curr_calories = 0

        for line in lines:
            if line.strip(SEPARATOR) == "":
                if curr_calories > 0:
                    calories.append(curr_calories)
                    curr_calories = 0
            else:
                curr_calories += int(line)

        top_three_elves = list(reversed(sorted(calories)))[:3]
        return sum(top_three_elves)


def main():
    print('Advent of Code - 2022 - Day 1 Solution')
    print(f'\tPart One Solution: {solutionPartOne()}')
    print(f'\tPart Two Solution: {solutionPartTwo()}')


if __name__ == "__main__":
    main()
