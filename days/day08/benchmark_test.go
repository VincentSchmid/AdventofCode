package day08

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day08(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(8))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day08(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(8))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}
