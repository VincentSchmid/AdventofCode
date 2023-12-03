package day15

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day15(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(15))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day15(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(15))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}
