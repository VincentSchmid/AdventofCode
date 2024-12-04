package utils

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Numeric interface {
	constraints.Signed | constraints.Float
}

func Abs[T Numeric](n T) T {
	if n < 0 {
		return -n
	}
	return n
}

func Max[T Numeric](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Min[T Numeric](a, b T) T {
	if a < b {
		return a
	}

	return b
}

func Clamp[T Numeric](low, n, high T) T {
	if low > high {
		panic(fmt.Errorf("IntClamp: low cannot be > high: %v > %v", low, high))
	}

	return Max(Min(n, high), low)
}

func ManhattanDistance[T Numeric](x1, y1, x2, y2 T) T {
	return Abs(x2-x1) + Abs(y2-y1)
}
