#!/usr/bin/env python3
"""
Advent Of Code[2022] - Day 2

Problem Link: https://adventofcode.com/2022/day/2
Difficulty: Easy
"""

INPUT_FILE = "input.txt"
SEPARATOR = "\n"

ROUND_VALUE = {"Rock": 1, "Paper": 2, "Scissors": 3}

ELF_PLAY = {
    "A": "Rock",
    "B": "Paper",
    "C": "Scissors",
}

WIN_YOUR_PLAY = {
    "X": ("Rock", "C"),
    "Y": ("Paper", "A"),
    "Z": ("Scissors", "B"),
}

DRAW_YOUR_PLAY = {
    "X": ("Rock", "A"),
    "Y": ("Paper", "B"),
    "Z": ("Scissors", "C"),
}

LOSE_YOUR_PLAY = {
    "X": ("Rock", "B"),
    "Y": ("Paper", "C"),
    "Z": ("Scissors", "A"),
}

WIN_INCREMENT = 6
DRAW_INCREMENT = 3
LOST_INCREMENT = 0

def solutionPartOne():
    with open(INPUT_FILE, "r") as file:
        content = file.read()
        lines = content.rstrip(SEPARATOR).split(SEPARATOR)

        GAME_SEPARATOR = " "
        final_score = 0
        for line in lines:
            elf_move, move = line.split(GAME_SEPARATOR)
            print(
                f"ELF Move: {elf_move}, Move: {move} Greater than: {WIN_YOUR_PLAY[move][1]}"
            )

            if WIN_YOUR_PLAY[move][1] == elf_move:
                final_score += WIN_INCREMENT + ROUND_VALUE[WIN_YOUR_PLAY[move][0]]
            elif WIN_YOUR_PLAY[move][0] == ELF_PLAY[elf_move]:
                final_score += DRAW_INCREMENT + ROUND_VALUE[WIN_YOUR_PLAY[move][0]]
            else:
                final_score += LOST_INCREMENT + ROUND_VALUE[WIN_YOUR_PLAY[move][0]]

        return final_score


# A Rock Y Rock
# B Paper X Rock
# C Scissors Z Rock


def main():
    with open(INPUT_FILE, "r") as file:
        content = file.read()
        lines = content.rstrip(SEPARATOR).split(SEPARATOR)

        GAME_SEPARATOR = " "
        final_score = 0
        for line in lines:
            elf_move, move = line.split(GAME_SEPARATOR)
            print(
                f"ELF Move: {elf_move}, Move: {move} Greater than: {WIN_YOUR_PLAY[move][1]}"
            )

            if move == "X": # NEED TO LOSE THE ROUND
                final_score += LOST_INCREMENT
            elif move == "Y": # NEED TO DRAW THE ROUND SO
                final_score += DRAW_INCREMENT + (ROUND_VALUE[ELF_PLAY[elf_move]])
            else: # NEED TO WIN THE ROUND SO
                final_score += WIN_INCREMENT + (ROUND_VALUE[ELF_PLAY[elf_move]])

        print(final_score)


if __name__ == "__main__":
    main()
