package day03

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
)

func Problem01(lines []string) int {
	var result int64
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)

	for _, line := range lines {
		results := re.FindAllStringSubmatch(line, -1)
		for _, match := range results {
			a, _ := strconv.ParseInt(match[1], 0, 64)
			b, _ := strconv.ParseInt(match[2], 0, 64)

			result += a * b
		}
	}

	return int(result)
}

func Problem02(lines []string) int {
	var result int64
	doRe := regexp.MustCompile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	enabled := true

	for _, line := range lines {

		dos := doRe.FindAllStringSubmatch(line, -1)

		for _, match := range dos {
			if match[0] == "don't()" {
				enabled = false
			}

			if match[0] == "do()" {
				enabled = true
			}

			if enabled && len(match) > 1 {
				a, _ := strconv.ParseInt(match[1], 0, 64)
				b, _ := strconv.ParseInt(match[2], 0, 64)
				result += a * b
			}
		}
	}

	return int(result)
}

func Run() {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(3, "input.txt"))
	fmt.Println(Problem01(lines))
	fmt.Println(Problem02(lines))
}
