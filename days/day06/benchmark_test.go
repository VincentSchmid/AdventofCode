package day06

import (
	"adventofcode/utils"
	"testing"
)

func BenchmarkProblem1Day06(b *testing.B) {
	lines, _ := utils.ReadInputFile("input.txt")

	for i := 0; i < b.N; i++ {
		Problem01(lines)
	}
}

func BenchmarkProblem2Day06(b *testing.B) {
	lines, _ := utils.ReadInputFile("input.txt")
	for i := 0; i < b.N; i++ {
		Problem02(lines)
	}
}
