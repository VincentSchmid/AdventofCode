package main

import (
	"fmt"
	"regexp"
	"strings"
	"strconv"
)

type Color string

const (
	Red   Color = "red"
	Green Color = "green"
	Blue  Color = "blue"
)

func main() {
	line := "Game 12: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	re := regexp.MustCompile(`Game (\d+):`)
    id := re.FindStringSubmatch(line)
    fmt.Println(id[1])

	allShowings := strings.SplitN(line, ":", 2)[1]
	showingsText := strings.Split(allShowings, ";")

	for _, showing := range showingsText {
        re := regexp.MustCompile(`(\d+)\s+(\w+)`)
        matches := re.FindStringSubmatch(showing)[1:]
		fmt.Println("Matches")
		fmt.Println(matches)

		showingGame := make(map[Color]int, 0)

		for i := 0; i < len(matches); i = i + 2 {
			num, _ := strconv.Atoi(matches[i])
			showingGame[Color(matches[i+1])] += num
		}
		fmt.Println(showingGame[Color("green")])
	}
}
