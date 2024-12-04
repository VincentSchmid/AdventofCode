package day02

import (
	"adventofcode/utils"
	"fmt"
	"strconv"
	"strings"
)

func parseLine(line string) ([]int, error) {
	parts := strings.Fields(line)
	numbers := make([]int, len(parts))

	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			return nil, fmt.Errorf("invalid number: %s", part)
		}
		numbers[i] = num
	}

	return numbers, nil
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isInOrder(numbers []int, maxDelta int) bool {
	ascending := numbers[0] < numbers[1]

	if numbers[0] == numbers[1] {
		return false
	}

	for i := 0; i < len(numbers)-1; i++ {
		delta := abs(numbers[i] - numbers[i+1])

		if delta > maxDelta {
			return false
		}

		if ascending {
			if numbers[i] >= numbers[i+1] {
				return false
			}
		}

		if !ascending {
			if numbers[i] <= numbers[i+1] {
				return false
			}
		}
	}

	return true
}

func removeAtIndex(arr []int, i int) []int {
	filtered := make([]int, 0, len(arr)-1)
	filtered = append(filtered, arr[:i]...)
	filtered = append(filtered, arr[i+1:]...)

	return filtered
}

func Problem01(lines []string) int {
	safeCount := 0

	for _, line := range lines {
		numbers, _ := parseLine(line)
		if isInOrder(numbers, 3) {
			safeCount++
		}
	}

	return int(safeCount)
}

func Problem02(lines []string) int {
	safeCount := 0

	for _, line := range lines {
		numbers, _ := parseLine(line)
		if isInOrder(numbers, 3) {
			safeCount++
		} else {
			for i := range numbers {
				numbersRemoved := removeAtIndex(numbers, i)

				if isInOrder(numbersRemoved, 3) {
					safeCount++
					break
				}
			}
		}
	}

	return int(safeCount)
}

func Run() {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(2, "input.txt"))
	fmt.Println(Problem01(lines))
	fmt.Println(Problem02(lines))
}
