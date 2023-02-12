package utils

import (
	"math/big"
	"sort"
)

func GetMedianOfBigIntSlice(numbers []*big.Int) *big.Int {
	sort.Slice(numbers, func(i, j int) bool {
		return numbers[i].Cmp(numbers[j]) == -1
	})

	length := len(numbers)
	median := big.NewInt(0)

	if length%2 == 0 {
		temp1 := big.NewInt(0).Set(numbers[length/2-1])
		temp2 := big.NewInt(0).Set(numbers[length/2])
		median = big.NewInt(0).Add(temp1, temp2)
		median = big.NewInt(0).Div(median, big.NewInt(2))
	} else {
		median = big.NewInt(0).Set(numbers[length/2])
	}

	return median
}

func WithinXPercent(x int, num1, num2 *big.Int) bool {
	diff := new(big.Int)
	diff.Sub(num1, num2)

	diff = diff.Abs(diff)

	percent := new(big.Int)
	percent.Mul(num1, big.NewInt(int64(x)))
	percent.Div(percent, big.NewInt(100))

	return diff.Cmp(percent) <= 0
}

func Max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}
