package day13

import (
	"adventofcode/utils"
	"testing"
)

func BenchmarkProblem1Day13(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(13))

	for i := 0; i < b.N; i++ {
		Problem01(lines)
	}
}

func BenchmarkProblem2Day13(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(13))
	for i := 0; i < b.N; i++ {
		Problem02(lines)
	}
}
