package utils

import (
	"bufio"
	"fmt"
	"os"
)

func GetInputFilePath(day int, filename string) string {
	return fmt.Sprintf("days/day%02d/%s", day, filename)
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
