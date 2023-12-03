#!/bin/bash

# Create or overwrite Go files for Advent of Code days 4 to 25
# Usage: ./create_day.sh

for DAY_NUM in {4..25}; do
    DAY_DIR="days/day$(printf "%02d" $DAY_NUM)"
    DAY_FILE="$DAY_DIR/day$(printf "%02d" $DAY_NUM).go"

    mkdir -p $DAY_DIR

    echo "package day$(printf "%02d" $DAY_NUM)

import (
    \"fmt\"
    \"adventofcode/utils\"
)

func Problem01(lines []string) int {
    return 0
}

func Problem02(lines []string) int {
    return 0
}

func Run() {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath($DAY_NUM))
    fmt.Println(Problem01(lines))
    fmt.Println(Problem02(lines))
}" > $DAY_FILE

    echo "Day $(printf "%02d" $DAY_NUM) Go file created or overwritten."
done