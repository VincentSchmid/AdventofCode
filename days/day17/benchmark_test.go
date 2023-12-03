package day17

import (
	"adventofcode/utils"
	"testing"
)

func BenchmarkProblem1Day17(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(17))

	for i := 0; i < b.N; i++ {
		Problem01(lines)
	}
}

func BenchmarkProblem2Day17(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(17))
	for i := 0; i < b.N; i++ {
		Problem02(lines)
	}
}
