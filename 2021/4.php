<?php

// PROBLEM: https://adventofcode.com/2021/day/4

$input_file = "input.txt";
$handle = fopen($input_file, "r");

/**
 * Checks if a Bingo board has won by having a complete row or column marked.
 *
 * @param array $board The 2D array representing the board. Each element is an array [$number, $is_marked].
 * @return bool True if the board has a winning row or column, false otherwise.
 */
function check_if_board_has_won($board) {
    $board_has_won = false;

    // Check for winning columns
    for ($j = 0; $j < count($board[0]); $j++) {
        $all_marked = true;
        for ($i = 0; $i < count($board); $i++) {
            [$number, $is_marked] = $board[$i][$j];
            if (!$is_marked) {
                $all_marked = false;
                break;
            }
        }

        if ($all_marked) {
            $board_has_won = true;
            break;
        }
    }

    if ($board_has_won) {
        return true;
    }

    for ($i = 0; $i < count($board); $i++) {
        $all_marked = true;
        for ($j = 0; $j < count($board[$i]); $j++) {
            [$number, $is_marked] = $board[$i][$j];
            if (!$is_marked) {
                $all_marked = false;
                break;
            }
        }

        if ($all_marked) {
            $board_has_won = true;
            break;
        }
    }

    return $board_has_won;
}

/**
 * Prints the entire board to the console for debugging/visualization.
 *
 * @param array $board The 2D array representing the board.
 * @return void
 */
function print_board($board) {
    for ($i = 0; $i < count($board); $i++) {
        print_board_row($board[$i]);
    }
}

/**
 * Helper function to print a single row of the board.
 *
 * @param array $board_row The 1D array representing a row of the board.
 * @return void
 */
function print_board_row($board_row) {
    printf("%s\n", implode(" ", array_map(function($value) {
        [$number, $is_marked] = $value;
        // Format: (Number, T/F for Marked)
        return sprintf("(%3d, %s)", $number, $is_marked ? "T" : "F");
    }, $board_row)));
}


/**
 * Marks a given number on the board as 'true' if it exists and is not already marked.
 * The board is passed by reference to modify it directly.
 *
 * @param array &$board The 2D array representing the board (passed by reference).
 * @param int $value The number to mark on the board.
 * @return void
 */
function mark_number_in_board(&$board, $value) {
    for ($i = 0; $i < count($board); $i++) {
        for ($j = 0; $j < count($board[$i]); $j++) {
            [$number, $is_marked] = $board[$i][$j];
            if ($number == $value && !$is_marked){
                $board[$i][$j] = [$number, true];
            }
        }
    }
}

/**
 * Calculates the sum of all unmarked numbers on the board.
 *
 * @param array $board The 2D array representing the board.
 * @return int The sum of all unmarked numbers.
 */
function sum_of_unmarked($board) {
    $sum = 0;
    for ($i = 0; $i < count($board); $i++) {
        for ($j = 0; $j < count($board[$i]); $j++) {
            [$number, $is_marked] = $board[$i][$j];
            if (!$is_marked) {
                $sum += $number;
            }
        }
    }
    return $sum;
}

/**
 * Reads the input file (via the provided file handle) to extract the sequence of draws
 * and the initial state of all Bingo boards.
 *
 * @param resource $handle The file resource handle opened for reading the input file.
 * @return array An array containing two elements: [$boards, $draws].
 * $boards is an array of 2D arrays (the boards).
 * $draws is an array of integers (the numbers to be drawn).
 */
function get_board_from_file($handle) {
    rewind($handle);

    $get_boards = false;
    $draws = [];
    $boards = [];
    $curr_board = [];
    while (($line = fgets($handle)) != false) {
        $line = trim($line);

        if ($line == "" && !$get_boards) {
            $get_boards = true;
            continue;
        }

        if (!$get_boards) {
            $draws = array_map(fn($value) => (int)$value, explode(",", $line));
        } else {
            if ($line == "") {
                array_push($boards, $curr_board);
                $curr_board = [];
            } else {
                array_push($curr_board, array_map(fn($value) => [(int)$value, false], preg_split('/\s+/', $line, -1, PREG_SPLIT_NO_EMPTY)));
            }
        }
    }

    // Add the last board if it wasn't followed by a trailing newline/empty line
    if (count($curr_board) > 0) {
        array_push($boards, $curr_board);
        $curr_board = [];
    }
    return [$boards, $draws];
}

/**
 * Solves the first part of the problem: finding the score of the *first* board to win.
 * The score is the sum of all unmarked numbers on the winning board multiplied by the last number drawn.
 *
 * @param resource $handle The file resource handle opened for reading the input file.
 * @return int The final score of the first winning board, or 0 if no board wins.
 */
function first_part_solution($handle) {
    [$boards, $draws] = get_board_from_file($handle);

    for ($i = 0; $i < count($draws); $i++) {
        for ($j = 0; $j < count($boards); $j++) {
            mark_number_in_board($boards[$j], $draws[$i]);
        }

        for ($k = 0; $k < count($boards); $k++) {
            if (check_if_board_has_won($boards[$k])) {
                return sum_of_unmarked($boards[$k]) * $draws[$i];
            }
        }
    }

    return 0;
}

/**
 * Solves the second part of the problem: finding the score of the *last* board to win.
 * The score is the sum of all unmarked numbers on the last winning board multiplied by the last number drawn.
 *
 * @param resource $handle The file resource handle opened for reading the input file.
 * @return int The final score of the last winning board, or -1 if no board wins (though likely guaranteed to win).
 */
function second_part_solution($handle) {
    [$boards, $draws] = get_board_from_file($handle);

    $last_board_to_win = -1;
    $boards_that_won = [];

    for ($i = 0; $i < count($draws); $i++) {
        for ($j = 0; $j < count($boards); $j++) {
            mark_number_in_board($boards[$j], $draws[$i]);
        }

        for ($k = 0; $k < count($boards); $k++) {
            if (check_if_board_has_won($boards[$k]) && !in_array($k, $boards_that_won)) {
                array_push($boards_that_won, $k);
                $last_board_to_win = sum_of_unmarked($boards[$k]) * $draws[$i];
            }
        }
    }

    return $last_board_to_win;
}

printf("Result \n\tFirst Part: %d\n\tSecond Part: %d\n", first_part_solution($handle), second_part_solution($handle));
