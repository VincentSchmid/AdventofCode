package day19

import (
	"adventofcode/utils"
	"testing"
)

func BenchmarkProblem1Day19(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(19))

	for i := 0; i < b.N; i++ {
		Problem01(lines)
	}
}

func BenchmarkProblem2Day19(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(19))
	for i := 0; i < b.N; i++ {
		Problem02(lines)
	}
}
