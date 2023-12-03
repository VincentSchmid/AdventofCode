package day12

import (
    "testing"
    "adventofcode/utils"
)

func BenchmarkProblem1Day12(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(12))

    for i := 0; i < b.N; i++ {
        Problem01(lines)
    }
}

func BenchmarkProblem2Day12(b *testing.B) {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(12))
    for i := 0; i < b.N; i++ {
        Problem02(lines)
    }
}
