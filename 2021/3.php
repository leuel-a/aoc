<?php
$input_file = "input.txt";
$lines = file($input_file, FILE_IGNORE_NEW_LINES | FILE_SKIP_EMPTY_LINES);

/**
 * Calculates the power consumption from binary numbers.
 *
 * @param string[] $lines Array of binary numbers as strings
 * @return int Product of gamma rate and epsilon rate
 */
function firstPartSolution($lines) {
        $gamma_rate = ""; $epsilon_rate = "";
        for ($j = 0; $j < strlen($lines[0]); $j++) {
                $one_count = 0; $zero_count = 0;
                for ($i = 0; $i < count($lines); $i++) {
                        if ($lines[$i][$j] == '0') $zero_count++; else $one_count++;
                }
                $gamma_rate   .= ($zero_count > $one_count) ? "0" : "1";
                $epsilon_rate .= ($zero_count > $one_count) ? "1" : "0";
        }
        return bindec($gamma_rate) * bindec($epsilon_rate);
}

/**
 * Calculates the life support rating from binary numbers.
 *
 * @param string[] $lines Array of binary numbers as strings
 * @return int Product of Oxygen Generator Rating and CO2 Scrubber Rating
 */
function secondPartSolution($lines) {
        // Oxygen Generator Rating
        $oxygen_rating = array_values($lines);
        $column = 0;
        while (count($oxygen_rating) > 1 && $column < strlen($oxygen_rating[0])) {
                $zeros_array = []; $ones_array = [];
                $one_count = 0; $zero_count = 0;
                for ($row = 0; $row < count($oxygen_rating); $row++) {
                        if ($oxygen_rating[$row][$column] == '0') {
                                $zero_count++; $zeros_array[] = $oxygen_rating[$row];
                        } else {
                                $one_count++; $ones_array[] = $oxygen_rating[$row];
                        }
                }
                $oxygen_rating = ($one_count >= $zero_count) ? array_values($ones_array) : array_values($zeros_array);
                $column++;
        }
        $oxygen_rating = array_map(fn($v) => bindec($v), $oxygen_rating);

        // CO2 Scrubber Rating
        $co2_scrubber_rating = array_values($lines);
        $column = 0;
        while (count($co2_scrubber_rating) > 1 && $column < strlen($co2_scrubber_rating[0])) {
                $zeros_array = []; $ones_array = [];
                $one_count = 0; $zero_count = 0;
                for ($row = 0; $row < count($co2_scrubber_rating); $row++) {
                        if ($co2_scrubber_rating[$row][$column] == '0') {
                                $zero_count++; $zeros_array[] = $co2_scrubber_rating[$row];
                        } else {
                                $one_count++; $ones_array[] = $co2_scrubber_rating[$row];
                        }
                }
                $co2_scrubber_rating = ($one_count >= $zero_count) ? array_values($zeros_array) : array_values($ones_array);
                $column++;
        }
        $co2_scrubber_rating = array_map(fn($v) => bindec($v), $co2_scrubber_rating);

        return $co2_scrubber_rating[0] * $oxygen_rating[0];
}

printf("Result \n\tFirst Part: %d\n\tSecond Part: %d\n", firstPartSolution($lines), secondPartSolution($lines));
