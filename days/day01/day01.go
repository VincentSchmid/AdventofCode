package day01

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func getNumbers(input string, separator string) ([]int64, error) {
	numbers := make([]int64, 0)

	pattern :=  fmt.Sprintf(`^\d+(%s\d+)*$`, separator)

	re := regexp.MustCompile(pattern)
	if !re.MatchString(input) {
		return nil, fmt.Errorf("input did not match required pattern")
	}

	numberStrings := strings.Split(input, separator)

	for _, number := range numberStrings {
		newNumber, _ := strconv.ParseInt(number, 0, 64)
		numbers = append(numbers, newNumber)
	}

	return numbers, nil
}

func Problem01(lines []string) int {
	lineCount := len(lines)
	left := make([]int64, lineCount)
	right := make([]int64, lineCount)

	for i, line := range lines {
		numbers, _ := getNumbers(line, "   ")
		left[i] = numbers[0]
		right[i] = numbers[1]
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	distances := make([]float64, lineCount)

	for i := 0; i < lineCount; i++ {
		distances[i] = math.Abs(float64(left[i] - right[i]))
	}

	var sum float64
	for _, num := range distances {
		sum += num
	}

	return int(sum)
}

func Problem02(lines []string) int {
	lineCount := len(lines)
	left := make([]int64, lineCount)
	right := make([]int64, lineCount)

	for i, line := range lines {
		numbers, _ := getNumbers(line, "   ")
		left[i] = numbers[0]
		right[i] = numbers[1]
	}

	sort.Slice(left, func(i, j int) bool {
		return left[i] < left[j]
	})

	sort.Slice(right, func(i, j int) bool {
		return right[i] < right[j]
	})

	var similarity int64
	checkIndex := 0

	for _, leftVal := range left {
		for leftVal > right[checkIndex] {
			checkIndex += 1
		}
		
		offset := 0
		for leftVal == right[checkIndex + offset] {
			similarity += leftVal
			offset += 1
		}
	}

	return int(similarity)
}

func Run() {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(1, "input.txt"))
	fmt.Println(Problem01(lines))
	fmt.Println(Problem02(lines))
}
