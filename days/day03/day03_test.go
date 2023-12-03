package day03

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
    arr := []string{".", "1", "2", "3", "."}

    result := getValue(&arr, 1)
    assert.Equal(t, result.Number, 123)
    assert.Equal(t, result.IndexesToJump, 3)

    result = getValue(&arr, 2)
    assert.Equal(t, result.Number, 123)
    assert.Equal(t, result.IndexesToJump, 2)

    result = getValue(&arr, 3)
    assert.Equal(t, result.Number, 123)
    assert.Equal(t, result.IndexesToJump, 1)
}
