<?php
$inputFile = "input.txt";
$lines = array_map('trim',file($inputFile));

/**
 * Calculates the final position product for the first part.
 *
 * @param string[] $lines Array of movement commands as "direction units"
 * @return int Absolute value of horizontal position multiplied by depth
 */
function firstPartSolution($lines) {
        $x = 0; $y = 0;
        foreach ($lines as $line) {
                [$direciton, $units] = explode(" ", $line);

                if (strcmp($direciton, "forward") == 0) {
                        $x = $x + (int)$units;
                } else if (strcmp($direciton, "down") == 0) {
                        $y = $y - (int)$units;
                } else if (strcmp($direciton, "up") == 0) {
                        $y = $y + (int)$units;
                }
        }
        return abs($y * $x);
}

/**
 * Calculates the final position product for the second part using aim.
 *
 * @param string[] $lines Array of movement commands as "direction units"
 * @return int Absolute value of horizontal position multiplied by depth
 */
function secondPartSolution($lines) {
        $x = 0; $y = 0; $aim = 0;
        foreach ($lines as $line) {
                [$direciton, $units] = explode(" ", $line);

                if (strcmp($direciton, "forward") == 0) {
                        $x = $x + (int)$units;
                        $y = $y + ($aim * (int)$units);
                } else if (strcmp($direciton, "down") == 0) {
                        $aim = $aim + (int)$units;
                } else if (strcmp($direciton, "up") == 0) {
                        $aim = $aim - (int)$units;
                }
        }
        return abs($y * $x);
}

printf("Result \n\tFirst Part: %d\n\tSecond Part: %d\n", firstPartSolution($lines), secondPartSolution($lines));
