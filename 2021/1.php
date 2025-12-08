<?php
// PROBLEM: https://adventofcode.com/2021/day/1

$input_file = 'input.txt';
$measurements = array_map('intval', file($input_file));

/**
 * Counts the number of times a measurement increases from the previous one.
 
 * @param int[] $measurements Array of integer measurements
 * @return int Number of increases
 */
function firstPartSolution($measurements) {
    $count = 0;
    for ($i = 1; $i < count($measurements); $i++) {
        if ($measurements[$i-1] < $measurements[$i]) {
            $count += 1;
        }
    }
    return $count;
}

/**
 * Counts the number of times the sum of a sliding window of size 3 increases.
 *
 * @param int[] $measurements Array of integer measurements
 * @return int Number of sliding window sum increases
 */
function secondPartSolution($measurements) {
    if (count($measurements) < 3) {
        return 0;
    }

    $count = 0;
    $window_sum = $measurements[0] + $measurements[1] + $measurements[2];
    for ($i = 1; $i < count($measurements) && $i + 2 < count($measurements); $i++) {
        $curr_window_sum = $window_sum - $measurements[$i - 1] + $measurements[$i + 2];

        if ($curr_window_sum > $window_sum) {
            $count++;
        }

        $window_sum = $curr_window_sum;
    }

    return $count;
}

printf(
    "Result\n\tFirst Part: %d\n\tSecond Part: %d\n",
    firstPartSolution($measurements),
    secondPartSolution($measurements)
);
