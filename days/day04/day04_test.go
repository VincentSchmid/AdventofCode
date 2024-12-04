package day04

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSmallDataSetP1(t *testing.T) {
	lines := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	result := Problem01(lines)
	fmt.Println(result)
	assert.Equal(t, 13, result)
}

func TestSmallDataSetP2(t *testing.T) {
	lines := []string{
		"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
	}

	result := Problem02(lines)
	fmt.Println(result)
	assert.Equal(t, 30, result)
}

func TestCards(t *testing.T) {
	cards := []ScratchCard{
		ScratchCardWithNMatches(3),
		ScratchCardWithNMatches(1),
		ScratchCardWithNMatches(2),
		ScratchCardWithNMatches(0),
		ScratchCardWithNMatches(0),
	}

	result := ProcessCards(&cards)

	assert.Equal(t, 18, result)
}

func TestFindMatchingNumbersInArray(t *testing.T) {
	arr1 := []int{1, 2, 4, 5, 6}
	arr2 := []int{1, 3, 4, 5}

	result := findMatchingNumbersInArray(arr1, arr2)

	assert.Equal(t, []int{1, 4, 5}, result)
}

func ScratchCardWithNMatches(n int) ScratchCard {
	return ScratchCard{id: 1, winningNumbers: []int{}, selectedNumbers: []int{}, matches: make([]int, n)}
}
