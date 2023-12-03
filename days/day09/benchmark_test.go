package day09

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day09(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(9))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day09(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(9))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}
