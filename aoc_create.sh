#!/bin/bash

# Create folders and files for Advent of Code days
# Usage: ./create_day.sh

for DAY_NUM in {1..25}; do
    DAY_DIR="days/day$(printf "%02d" $DAY_NUM)"

    mkdir -p $DAY_DIR
    touch $DAY_DIR/day$(printf "%02d" $DAY_NUM).go
    touch $DAY_DIR/day$(printf "%02d" $DAY_NUM)_test.go
    touch $DAY_DIR/benchmark_test.go
    touch $DAY_DIR/input.txt

    echo "package day$(printf "%02d" $DAY_NUM)

func Run() {
    // Day $(printf "%02d" $DAY_NUM) implementation
}" > $DAY_DIR/day$(printf "%02d" $DAY_NUM).go

    echo "package day$(printf "%02d" $DAY_NUM)

import (
    \"testing\"
    \"github.com/stretchr/testify/assert\"
)

func TestDay$(printf "%02d" $DAY_NUM)(t *testing.T) {
    assert.True(t, true) // Replace with actual tests
}" > $DAY_DIR/day$(printf "%02d" $DAY_NUM)_test.go

    echo "package day$(printf "%02d" $DAY_NUM)

import \"testing\"

func BenchmarkProblem1Day$(printf "%02d" $DAY_NUM)(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Benchmark code for Problem 1
    }
}

func BenchmarkProblem2Day$(printf "%02d" $DAY_NUM)(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Benchmark code for Problem 2
    }
}" > $DAY_DIR/benchmark_test.go

    echo "Day $(printf "%02d" $DAY_NUM) structure with input file and benchmarks for Problem 1 and 2 created."
done
