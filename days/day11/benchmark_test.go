package day11

import (
	"adventofcode/utils"
	"testing"
)

func BenchmarkProblem1Day11(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(11))

	for i := 0; i < b.N; i++ {
		Problem01(lines)
	}
}

func BenchmarkProblem2Day11(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(11))
	for i := 0; i < b.N; i++ {
		Problem02(lines)
	}
}
