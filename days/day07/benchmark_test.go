package day07

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day07(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(7))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day07(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(7))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}
