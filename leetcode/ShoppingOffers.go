package leetcode

import "math"

func isClean(needs []int) bool {
	for i := 0; i < len(needs); i++ {
		if needs[i] != 0 {
			return false
		}
	}

	return true
}

func remainneeds(needs, price []int) int {
	total := 0
	for i := 0; i < len(needs); i++ {
		total += needs[i] * price[i]
	}
	return total
}

func minpays(pays []int) int {
	min := math.MaxInt32
	for i := 0; i < len(pays); i++ {
		if pays[i] < min {
			min = pays[i]
		}
	}
	return min
}

func minusneeds(needs, special []int) []int {

	for i := 0; i < len(needs); i++ {
		needs[i] = needs[i] - special[i]
	}
	return needs
}

func canSpecial(needs, special []int) bool {
	for i := range needs {
		if needs[i] < special[i] {
			return false
		}
	}
	return true
}

func shoppingOffers(price []int, special [][]int, needs []int) int {

	if isClean(needs) {
		return 0
	}

	pays := []int{}

	for i := 0; i < len(special); i++ {
		if canSpecial(needs, special[i]) {
			nowneed := make([]int, len(needs))
			copy(nowneed, needs)

			for j := 0; j < len(nowneed); j++ {
				nowneed[j] -= special[i][j]
			}

			pay := special[i][len(nowneed)] + shoppingOffers(price, special, nowneed)
			pays = append(pays, pay)
		}
	}

	pays = append(pays, remainneeds(needs, price))

	return minpays(pays)
}
