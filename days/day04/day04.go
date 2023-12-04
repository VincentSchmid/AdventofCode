package day04

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	SCRATCH_CARD_REGEX = `Card\s+(\d+):\s+((\d+\s+)+)\|\s+((\d+\s*)+)`
)

func findMatchingNumbersInArray(arr1 []int, arr2 []int) []int {
	matches := make([]int, 0)
	sort.Sort(sort.IntSlice(arr1))
	sort.Sort(sort.IntSlice(arr2))

	i, j := 0, 0
	for i < len(arr1) && j < len(arr2) {
		if arr1[i] == arr2[j] {
			matches = append(matches, arr1[i])
			i++
			j++
		} else if arr1[i] < arr2[j] {
			i++
		} else {
			j++
		}
	}

	return matches
}

type ScratchCard struct {
	id              int
	winningNumbers  []int
	selectedNumbers []int
}

func NewScratchCard(scratchCardText string) *ScratchCard {
	scRe := regexp.MustCompile(SCRATCH_CARD_REGEX)
	result := scRe.FindStringSubmatch(scratchCardText)
	id, _ := strconv.Atoi(result[1])
	winningNumbersText := strings.Fields(result[2])
	selectedNumberText := strings.Fields(result[4])

	winningNumbers := utils.StrArrToIntArr(&winningNumbersText)
	selectedNumber := utils.StrArrToIntArr(&selectedNumberText)

	return &ScratchCard{
		id:              id,
		winningNumbers:  winningNumbers,
		selectedNumbers: selectedNumber,
	}
}

func (s *ScratchCard) GetPoints() int {
	matches := findMatchingNumbersInArray(s.selectedNumbers, s.winningNumbers)
	return int(math.Pow(2, float64(len(matches)-1)))
}

func Problem01(lines []string) int {
	sum := 0
	for _, line := range lines {
		card := *NewScratchCard(line)
		sum += card.GetPoints()
	}

	return sum
}

func Problem02(lines []string) int {
	return 0
}

func Run() {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(4))
	fmt.Println(Problem01(lines))
	fmt.Println(Problem02(lines))
}
