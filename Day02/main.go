package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Color string

type GameCubes struct {
	Red   int
	Green int
	Blue  int
}

type Game struct {
	id       int
	showings []map[Color]int
}

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

var (
	cubesInGame = GameCubes{Red: 12, Green: 13, Blue: 14}
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

func getMaxValue(arr []int) int {
    if len(arr) == 0 {
        return 0
    }

    max := arr[0]

    for _, num := range arr {
        if num > max {
            max = num
        }
    }

    return max
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func GameFromLine(line string) Game {
	reId := regexp.MustCompile(`Game (\d+):`)
	reShowing := regexp.MustCompile(`(\d+)\s+(\w+)`)

	id_matches := reId.FindStringSubmatch(line)
	id, _ := strconv.Atoi(id_matches[1])
	allShowings := strings.SplitN(line, ":", 2)[1]
	showingsText := strings.Split(allShowings, ";")

	allTheShowings := []map[Color]int{}

	for _, showing := range showingsText {
		matches := reShowing.FindAllStringSubmatch(showing, -1)

		showingGame := make(map[Color]int, 0)

		for _, col := range matches {
			num, _ := strconv.Atoi(col[1])
			showingGame[Color(col[2])] += num
		}

		allTheShowings = append(allTheShowings, showingGame)
	}

	return Game{id: id, showings: allTheShowings}
}

func isGameValid(game Game) bool {
	for _, showing := range game.showings {
		if showing[Red] > cubesInGame.Red || showing[Green] > cubesInGame.Green || showing[Blue] > cubesInGame.Blue {
			return false
		}
	}

	return true
}

func cubePower(game Game) int {
    maxRed := 0
    maxGreen := 0
    maxBlue := 0
    for _, showing := range game.showings {
        maxRed = max(maxRed, showing[Red])
        maxGreen = max(maxGreen, showing[Green])
        maxBlue = max(maxBlue, showing[Blue])
    }

    return maxRed * maxGreen * maxBlue
}

func main() {
	fmt.Println("Advent of Code 2022 - Day 2")
	lines, _ := ReadInputFile("./input.txt")

	games := make([]Game, 0)

	for _, line := range lines {
		games = append(games, GameFromLine(line))
	}

    // Problem 1
	sum := 0

	for _, game := range games {
		if isGameValid(game) {
			sum += game.id
		}
	}

	fmt.Println(sum)

    // Problem 2
    sum = 0

	for _, game := range games {
		sum += cubePower(game)
	}

	fmt.Println(sum)
}
