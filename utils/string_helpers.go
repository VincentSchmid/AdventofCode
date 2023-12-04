package utils

import (
	"strconv"
)

func IsDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func StrArrToIntArr(arr *[]string) []int {
	intArr := make([]int, len(*arr))
	for i, strVal := range *arr {
		val, _ := strconv.Atoi(strVal)
		intArr[i] = val
	}

	return intArr
}
