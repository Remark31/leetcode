package leetcode

import()

func isPalindrome(x int) bool {
    if x < 0 {
		return false
	}

	ans := 0
	
	t := x

	for t > 0 {
		ans *= 10
		d := t % 10
		ans += d
		t /= 10
	}

	return ans - x == 0


}