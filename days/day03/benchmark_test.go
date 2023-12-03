package day03

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day03(b *testing.B) {
    lines, _ := utils.ReadInputFile("input.txt")

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day03(b *testing.B) {
    for i := 0; i < b.N; i++ {
        // Benchmark code for Problem 2
    }
}
