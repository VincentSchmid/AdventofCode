package utils

import (
	"bufio"
	"os"
	"fmt"
)

func GetInputFilePath(day int) string {
    return fmt.Sprintf("days/day%02d/input.txt", day)
}

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

