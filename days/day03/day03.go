package day03

import (
	"adventofcode/utils"
	"fmt"
	"regexp"
	"strconv"
)

const (
	SEPARATOR    = '.'
	GEAR         = '*'
	SYMBOL_REGEX = `[^A-Za-z0-9.]`
)

type FoundNumber struct {
	Number        int
	IndexesToJump int
}

type Vector struct {
	x int
	y int
}

func getValue(line string, index int) (int, int) {
	digits := string(line[index])
	numDigits := 1
	i := index - 1

	for i >= 0 && IsDigit(line[i]) {
		digits = string(line[i]) + digits
		i--
		numDigits++
	}

	i = index + 1
	for i < len(line) && IsDigit(line[i]) {
		digits += string(line[i])
		i++
		numDigits++
	}

	number, _ := strconv.Atoi(digits)

	return number, numDigits
}

func ValueTouchesSymbol(valueArea []string) bool {
	re := regexp.MustCompile(SYMBOL_REGEX)

	for _, s := range valueArea {
		if re.MatchString(s) {
			return true
		}
	}

	return false
}

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func ExpandArea(topLeft Vector, bottomRight Vector, maxTopLeft Vector, maxBottomRight Vector, extend int) (Vector, Vector) {
	newTopLeft := Vector{x: utils.IntMax(maxTopLeft.x, topLeft.x-extend), y: utils.IntMax(maxTopLeft.y, topLeft.y-extend)}
	newBottomRight := Vector{x: utils.IntMin(maxBottomRight.x, bottomRight.x+extend+1), y: utils.IntMin(maxBottomRight.y, bottomRight.y+extend+1)}

	return newTopLeft, newBottomRight
}

func getValueArea(arr *[]string, x int, y int, width int, maxTopLeft Vector, maxBottomRight Vector, extend int) []string {
	newTopLeft, newBottomRight := ExpandArea(Vector{x: x, y: y}, Vector{x: x + width, y: y}, maxTopLeft, maxBottomRight, extend)

	valueArea := make([]string, newBottomRight.y-newTopLeft.y)
	copy(valueArea, (*arr)[newTopLeft.y:newBottomRight.y])

	for i, row := range valueArea {
		valueArea[i] = row[newTopLeft.x:newBottomRight.x]
	}

	return valueArea
}

func Problem01(lines []string) int {
	sum := 0
	lineLength := len(lines[0])
	maxTopLeft := Vector{x: 0, y: 0}
	maxBottomRight := Vector{x: len(lines[0]), y: len(lines)}

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if IsDigit(line[x]) {
				val, numDigits := getValue(line, x)
				valueArea := getValueArea(&lines, x, y, numDigits-1, maxTopLeft, maxBottomRight, 1)

				if ValueTouchesSymbol(valueArea) {
					sum += val
				}

				x += utils.IntMin(numDigits, lineLength)
			}
		}
	}

	return sum
}

func Problem02(lines []string) int {
	sum := 0
	maxTopLeft := Vector{x: 0, y: 0}
	maxBottomRight := Vector{x: len(lines[0]), y: len(lines)}

	for y, line := range lines {
		for x := 0; x < len(line); x++ {
			if line[x] == GEAR {
				valueArea := getValueArea(&lines, x, y, 0, maxTopLeft, maxBottomRight, 1)
				fmt.Println(valueArea)
			}
		}
	}

	return sum
}

func Run() {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(3))
	fmt.Println(Problem01(lines))
}
