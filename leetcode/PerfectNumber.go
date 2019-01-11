package leetcode

import (
	"math"
)

func checkPerfectNumber(num int) bool {
	if num <= 1 {
		return false
	}
	l := math.Sqrt(float64(num))
	il := int(l)
	sum := 1
	for i := 2; i <= il; i++ {
		if num%i == 0 {
			sum += i + num/i
		}
		if sum > num {
			return false
		}
	}
	return sum == num
}
