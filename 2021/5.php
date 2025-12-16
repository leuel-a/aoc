<?php

// PROBLEM: https://adventofcode.com/2021/day/5

$input_file = "input.txt";
$lines = file($input_file, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);

$ROW_LENGTH = 1000;
$COLUMN_LENGTH = 1000;

/**
 * Initializes a 2D grid with the given dimensions.
 *
 * Each cell in the grid is initialized to 0.
 *
 * @param int $rows    Number of rows in the grid.
 * @param int $columns Number of columns in the grid.
 *
 * @return array<int, array<int, int>> A 2D grid of integers.
 */
function initialize_grid($rows = 10, $columns = 10) {
        $grid = [];

        for ($i = 0; $i < $rows; $i++) {
                $grid[$i] = [];
                for ($j = 0; $j < $columns; $j++) {
                        $grid[$i][$j] = 0;
                }
        }
        return $grid;
}

/**
 * Solves the first part of the problem.
 *
 * Processes only horizontal and vertical line segments and counts
 * the number of grid points where at least two lines overlap.
 *
 * @param array<int, string> $lines Input lines in the format "x1,y1 -> x2,y2".
 *
 * @return int Number of overlapping grid points (> 1).
 */
function firstPartSolution($lines) {
        global $ROW_LENGTH, $COLUMN_LENGTH;

        $pattern = '/(\d+),(\d+) -> (\d+),(\d+)/';
        $grid = initialize_grid($ROW_LENGTH, $COLUMN_LENGTH);

        foreach ($lines as $line) {
                $first = ["x" => 0, "y" => 0];
                $second = ["x" => 0, "y" => 0];

                if (preg_match($pattern, $line, $matches)) {
                        [$first["x"], $first["y"], $second["x"], $second["y"]] =
                                array_map('intval', array_slice($matches, 1));
                }

                // Ignore single-point lines
                if ($first["x"] == $second["x"] && $first["y"] == $second["y"]) {
                        continue;
                }

                // Horizontal line
                if ($first["y"] == $second["y"]) {
                        $difference = abs($first["x"] - $second["x"]);
                        $initial = min($first["x"], $second["x"]);

                        for ($i = 0; $i <= $difference; $i++) {
                                $grid[$first["y"]][$initial + $i] += 1;
                        }
                }
                // Vertical line
                else if ($first["x"] == $second["x"]) {
                        $difference = abs($first["y"] - $second["y"]);
                        $initial = min($first["y"], $second["y"]);

                        for ($j = 0; $j <= $difference; $j++) {
                                $grid[$initial + $j][$first["x"]] += 1;
                        }
                }
        }

        // Count overlaps
        $result = 0;
        for ($i = 0; $i < count($grid); $i++) {
                for ($j = 0; $j < count($grid[$i]); $j++) {
                        if ($grid[$i][$j] > 1) {
                                $result++;
                        }
                }
        }
        return $result;
}

/**
 * Solves the second part of the problem.
 *
 * Processes horizontal, vertical, and diagonal (45-degree) line segments
 * and counts the number of grid points where at least two lines overlap.
 *
 * @param array<int, string> $lines Input lines in the format "x1,y1 -> x2,y2".
 *
 * @return int Number of overlapping grid points (> 1).
 */
function secondPartSolution($lines) {
        global $ROW_LENGTH, $COLUMN_LENGTH;

        $pattern = '/(\d+),(\d+) -> (\d+),(\d+)/';
        $grid = initialize_grid($ROW_LENGTH, $COLUMN_LENGTH);

        foreach ($lines as $line) {
                $first = ["x" => 0, "y" => 0];
                $second = ["x" => 0, "y" => 0];

                if (preg_match($pattern, $line, $matches)) {
                        [$first["x"], $first["y"], $second["x"], $second["y"]] =
                                array_map('intval', array_slice($matches, 1));
                }

                // Horizontal
                if ($first["y"] == $second["y"]) {
                        $difference = abs($first["x"] - $second["x"]);
                        $initial = min($first["x"], $second["x"]);

                        for ($i = 0; $i <= $difference; $i++) {
                                $grid[$first["y"]][$initial + $i] += 1;
                        }
                }
                // Vertical
                else if ($first["x"] == $second["x"]) {
                        $difference = abs($first["y"] - $second["y"]);
                        $initial = min($first["y"], $second["y"]);

                        for ($j = 0; $j <= $difference; $j++) {
                                $grid[$initial + $j][$first["x"]] += 1;
                        }
                }
                // Diagonal (45 degrees)
                else if (abs($first["x"] - $second["x"]) == abs($first["y"] - $second["y"])) {
                        $difference = abs($first["x"] - $second["x"]);

                        for ($i = 0; $i <= $difference; $i++) {
                                $x = $first["x"] + ($second["x"] > $first["x"] ? $i : -$i);
                                $y = $first["y"] + ($second["y"] > $first["y"] ? $i : -$i);
                                $grid[$y][$x] += 1;
                        }
                }
        }

        // Count overlaps
        $result = 0;
        for ($i = 0; $i < count($grid); $i++) {
                for ($j = 0; $j < count($grid[$i]); $j++) {
                        if ($grid[$i][$j] > 1) {
                                $result++;
                        }
                }
        }
        return $result;
}

/**
 * Prints the entire grid to stdout.
 *
 * @param array<int, array<int, int>> $grid The grid to print.
 *
 * @return void
 */
function print_grid($grid) {
        for ($i = 0; $i < count($grid); $i++) {
                print_grid_row($grid[$i]);
        }
}

/**
 * Prints a single row of the grid.
 *
 * Zero values are displayed as dots (.), non-zero values as numbers.
 *
 * @param array<int, int> $grid_row A single row from the grid.
 *
 * @return void
 */
function print_grid_row($grid_row) {
        printf(
                "%s\n",
                implode(
                        "",
                        array_map(function ($value) {
                                return $value === 0
                                        ? '.'
                                        : str_pad((string)$value, 1, " ", STR_PAD_BOTH);
                        }, $grid_row)
                )
        );
}

printf(
        "Result:\n\tFirst Part: %d\n\tSecond Part: %d\n",
        firstPartSolution($lines),
        secondPartSolution($lines)
);
