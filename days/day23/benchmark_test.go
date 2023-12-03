package day23

import (
	"adventofcode/utils"
	"testing"
)

func BenchmarkProblem1Day23(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(23))

	for i := 0; i < b.N; i++ {
		Problem01(lines)
	}
}

func BenchmarkProblem2Day23(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(23))
	for i := 0; i < b.N; i++ {
		Problem02(lines)
	}
}
