package day02

import (
	"adventofcode/utils"
	"fmt"
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
		maxRed = utils.Max(maxRed, showing[Red])
		maxGreen = utils.Max(maxGreen, showing[Green])
		maxBlue = utils.Max(maxBlue, showing[Blue])
	}

	return maxRed * maxGreen * maxBlue
}

func Run() {
	fmt.Println("Advent of Code 2022 - Day 2")
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(2))

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
