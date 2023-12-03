#!/bin/bash

# Create or overwrite Go files and Benchmark test files for Advent of Code days 4 to 25
# Usage: ./create_day.sh

for DAY_NUM in {4..25}; do
    DAY_DIR="days/day$(printf "%02d" $DAY_NUM)"
    DAY_FILE="$DAY_DIR/day$(printf "%02d" $DAY_NUM).go"
    BENCHMARK_FILE="$DAY_DIR/benchmark_test.go"

    mkdir -p $DAY_DIR

    # Create or overwrite dayXX.go
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

    # Create or overwrite benchmark_test.go
    echo "package day$(printf "%02d" $DAY_NUM)

import (
    \"testing\"
    \"adventofcode/utils\"
)

func BenchmarkProblem1Day$(printf "%02d" $DAY_NUM)(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath($DAY_NUM))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day$(printf "%02d" $DAY_NUM)(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath($DAY_NUM))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}" > $BENCHMARK_FILE

    echo "Day $(printf "%02d" $DAY_NUM) Go file and Benchmark test file created or overwritten."
done
