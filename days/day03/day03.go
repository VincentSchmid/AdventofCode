package day03

import (
    "fmt"
    "strconv"
    "adventofcode/utils"
)

const (
    SEPARATOR = "."
)

type FoundNumber struct {
	Number int
	IndexesToJump int
}

func getValue(arr *[]string, index int) FoundNumber {
    i := index
    digits := (*arr)[i]
	offset := 1
    for (*arr)[i] == SEPARATOR {
        i--
        digits = (*arr)[i] + digits
    }

	i = index

	for (*arr)[i] == SEPARATOR {
		i++
		offset++
		digits += (*arr)[i]
	}

	number, _ := strconv.Atoi(digits)

	return FoundNumber{Number: number, IndexesToJump: offset}
}

func Problem01(lines []string) int {
    return 0
}

func Run() {
    lines, _ := utils.ReadInputFile(utils.GetInputFilePath(3))
    fmt.Println(Problem01(lines))
}
