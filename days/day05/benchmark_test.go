package day05

import (
	"adventofcode/utils"
	"testing"
)

func BenchmarkProblem1Day05(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(5))

	for i := 0; i < b.N; i++ {
		Problem01(lines)
	}
}

func BenchmarkProblem2Day05(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(5))
	for i := 0; i < b.N; i++ {
		Problem02(lines)
	}
}
