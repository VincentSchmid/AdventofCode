package day16

import (
	"adventofcode/utils"
	"testing"
)

func BenchmarkProblem1Day16(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(16))

	for i := 0; i < b.N; i++ {
		Problem01(lines)
	}
}

func BenchmarkProblem2Day16(b *testing.B) {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(16))
	for i := 0; i < b.N; i++ {
		Problem02(lines)
	}
}
