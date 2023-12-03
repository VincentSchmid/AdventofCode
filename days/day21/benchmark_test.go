package day21

import (
	"adventofcode/utils"
	"testing"
)

func BenchmarkProblem1Day21(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(21))

	for i := 0; i < b.N; i++ {
		Problem01(lines)
	}
}

func BenchmarkProblem2Day21(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(21))
	for i := 0; i < b.N; i++ {
		Problem02(lines)
	}
}
