package day22

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day22(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(22))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day22(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(22))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}
