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
	sort.Ints(arr1)
	sort.Ints(arr2)

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
	matches         []int
}

func NewScratchCard(scratchCardText string) *ScratchCard {
	scRe := regexp.MustCompile(SCRATCH_CARD_REGEX)
	result := scRe.FindStringSubmatch(scratchCardText)
	id, _ := strconv.Atoi(result[1])
	winningNumbersText := strings.Fields(result[2])
	selectedNumberText := strings.Fields(result[4])

	winningNumbers := utils.StrArrToIntArr(&winningNumbersText)
	selectedNumber := utils.StrArrToIntArr(&selectedNumberText)
	matches := findMatchingNumbersInArray(selectedNumber, winningNumbers)

	return &ScratchCard{
		id:              id,
		winningNumbers:  winningNumbers,
		selectedNumbers: selectedNumber,
		matches:         matches,
	}
}

func (s *ScratchCard) GetPoints() int {
	return int(math.Pow(2, float64(len(s.matches)-1)))
}

func ProcessCards(cards *[]ScratchCard) int {
	cardCount := make([]int, len(*cards))
	sum := len(cardCount)
	for i := len(*cards) - 1; i >= 0; i-- {
		card := (*cards)[i]
		cardCount[i] += len(card.matches)

		for j := range card.matches {
			if i+j+1 < 0 {
				break
			}

			cardCount[i] += cardCount[i+j+1]
		}

		sum += cardCount[i]
	}

	return sum
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

	cards := make([]ScratchCard, len(lines))
	for i := len(lines) - 1; i >= 0; i-- {
		cards[i] = *NewScratchCard(lines[i])
	}

	result := ProcessCards(&cards)

	return result
}

func Run() {
	lines, _ := utils.ReadInputFile(utils.GetInputFilePath(4, "input.txt"))
	fmt.Println(Problem01(lines))
	fmt.Println(Problem02(lines))
}
