package day18

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day18(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(18))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day18(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(18))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}
