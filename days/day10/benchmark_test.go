package day10

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day10(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(10))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day10(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(10))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}
