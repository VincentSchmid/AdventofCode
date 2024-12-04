package day03

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Problem01(lines []string) int {
	var result int64
	re, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)

	for _, line := range lines {
		results := re.FindAllStringSubmatch(line, -1)
		for _, match := range(results) {
			a, _ := strconv.ParseInt(match[1], 0, 64)
			b, _ := strconv.ParseInt(match[2], 0, 64)

			result += a * b
		}
	}

	return int(result)
}

func Problem02(lines []string) int {
	var result int64
	doRe, _ := regexp.Compile(`do\(\)|don't\(\)|mul\((\d+),(\d+)\)`)
	re, _ := regexp.Compile(`mul\((\d+),(\d+)\)`)
	calculate := true

	for _, line := range lines {
		
		dos := doRe.FindAllString(line, -1)

		for _, match := range(dos) {
			if match == "don't()" {
				calculate = false
			}

			if match == "do()" {
				calculate = true
			}

			if calculate && strings.Contains(match, "mul") {
				match2 := re.FindAllStringSubmatch(match, -1)
				a, _ := strconv.ParseInt(match2[0][1], 0, 64)
				b, _ := strconv.ParseInt(match2[0][2], 0, 64)
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
