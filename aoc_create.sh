#!/bin/bash

# Create folders and files for Advent of Code days
# Usage: ./aoc_create.sh

for DAY_NUM in {1..25}; do
    DAY_DIR="days/day$(printf "%02d" $DAY_NUM)"
    DAY_FILE="$DAY_DIR/day$(printf "%02d" $DAY_NUM).go"
    TEST_FILE="$DAY_DIR/day$(printf "%02d" $DAY_NUM)_test.go"
    BENCHMARK_FILE="$DAY_DIR/benchmark_test.go"

    mkdir -p $DAY_DIR
    touch $DAY_DIR/input.txt

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

    # Create or overwrite dayXX_test.go
    echo "package day$(printf "%02d" $DAY_NUM)

import (
    \"testing\"
    \"github.com/stretchr/testify/assert\"
)

func TestDay$(printf "%02d" $DAY_NUM)(t *testing.T) {
    assert.True(t, true) // Replace with actual tests
}" > $TEST_FILE

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

    echo "Day $(printf "%02d" $DAY_NUM) Go file, test file, and benchmark test file created or overwritten."
done
