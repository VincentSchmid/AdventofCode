package day04

import (
	"adventofcode/utils"
	"testing"
)

func BenchmarkProblem1Day04(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(4))

	for i := 0; i < b.N; i++ {
		Problem01(lines)
	}
}

func BenchmarkProblem2Day04(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(4))
	for i := 0; i < b.N; i++ {
		Problem02(lines)
	}
}
