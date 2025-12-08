<?php
$inputFile = "input.txt";
$lines = array_map('trim',file($inputFile));

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
