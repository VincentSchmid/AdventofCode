package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var (
	digits = map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
)

func ReadInputFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func firstSpelledNumber(line string) string {
	for i := 0; i < len(line); i++ {
		token := ""
		for j := 0; j < 5; j++ {
			if i+j >= len(line) {
				break
			}

			token = line[i : i+j+1]

			for k, v := range digits {
				if k == token {
					return v
				}
			}
		}
	}

	return ""
}

func lastSpelledNumber(line string) string {
	for i := len(line) - 1; i >= 0; i-- {
		token := ""
		for j := 0; j < 5; j++ {
			if i-j < 0 {
				break
			}

			token = string(line[i-j]) + token

			for k, v := range digits {
				if k == token {
					return v
				}
			}
		}
	}

	return ""
}

func GetNumbers(line string) []string {
	re := regexp.MustCompile("[0-9]")

	return re.FindAllString(line, -1)
}

func indexOfFirstDigit(s string) int {
	for i, r := range s {
		if unicode.IsDigit(r) {
			return i
		}
	}
	return -1
}

func indexOfLastDigit(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(s[i])) {
			return i
		}
	}
	return -1
}

func Problem(lines []string, getNumbersInString func(string) []string) int {
	sum := 0
	var num int
	for _, line := range lines {
		onlyDigits := getNumbersInString(line)

		if len(onlyDigits) != 0 {
			num, _ = strconv.Atoi(onlyDigits[0] + onlyDigits[len(onlyDigits)-1])
		}
		sum += num
	}

	return sum
}

func main() {
	fmt.Println("Advent of Code 2022 - Day 1")

	lines, _ := ReadInputFile("input.txt")

	// Problem 1
	fmt.Println("Problem 1")
	fmt.Printf("%d\n", Problem(lines, GetNumbers))

	// Problem 2
	problem2Func := func(line string) []string {
		firstNumber := "0"
		lastNumber := "0"
		firstDigitIndex := indexOfFirstDigit(line)

		if firstDigitIndex == -1 {
			firstDigitIndex = len(line)
		} else {
			firstNumber = string(line[indexOfFirstDigit(line)])
		}

		spelledNumber := firstSpelledNumber(line[:firstDigitIndex])

		if spelledNumber != "" {
			firstNumber = spelledNumber
		}

		lastDigitIndex := indexOfLastDigit(line)
		if lastDigitIndex == -1 {
			lastDigitIndex = 0
		} else {
			lastNumber = string(line[indexOfLastDigit(line)])
		}

		spelledNumber = lastSpelledNumber(line[lastDigitIndex:])

		if spelledNumber != "" {
			lastNumber = spelledNumber
		}

		return []string{firstNumber, lastNumber}
	}

	fmt.Println("Problem 2")
	fmt.Printf("%d\n", Problem(lines, problem2Func))
}
