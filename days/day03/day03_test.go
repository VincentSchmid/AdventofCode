package day03

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
    line := ".123."

    result, numDigits := getValue(line, 1)
    assert.Equal(t, 123, result)
    assert.Equal(t, 3, numDigits)

    line = ".123"

    result, numDigits = getValue(line, 1)
    assert.Equal(t, 123, result)
    assert.Equal(t, 3, numDigits)

    line = ".123*"

    result, numDigits = getValue(line, 1)
    assert.Equal(t, 123, result)
    assert.Equal(t, 3, numDigits)
}

func TestExpandArea(t *testing.T) {
    topLeft := Vector{x: 1, y: 1}
    bottomRight := Vector{x: 3, y: 3}
    maxTopLeft := Vector{x: 0, y: 0}
    maxBottomRight := Vector{x: 4, y: 4}

    newTopLeft, newBottomRight := ExpandArea(topLeft, bottomRight, maxTopLeft, maxBottomRight, 1)

    assert.Equal(t, Vector{x: 0, y: 0}, newTopLeft)
    assert.Equal(t, Vector{x: 4, y: 4}, newBottomRight)

    maxTopLeft = Vector{x: 2, y: 2}
    maxBottomRight = Vector{x: 3, y: 3}

    newTopLeft, newBottomRight = ExpandArea(topLeft, bottomRight, maxTopLeft, maxBottomRight, 1)

    assert.Equal(t, Vector{x: 2, y: 2}, newTopLeft)
    assert.Equal(t, Vector{x: 3, y: 3}, newBottomRight)
}

func TestBug(t *testing.T) {
    lines := []string{
        "..515",
    }

    result := Problem01(lines)
    assert.Equal(t, 0, result)

    lines = []string{
        ".-515",
    }

    result = Problem01(lines)
    assert.Equal(t, 515, result)
}

func TestSmallInput(t *testing.T) {
    lines := []string {
        "467..114..",
        "...*......",
        "..35..633.",
        "......#...",
        "617*......",
        ".....+.58.",
        "..592.....",
        "......755.",
        "...$.*....",
        ".664.598..",
    }

    result := Problem01(lines)
    assert.Equal(t, 4361, result)
}
