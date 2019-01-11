package leetcode

func longestValidParentheses(s string) int {
	d := make([]int, len(s))
	left := 0
	max := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left += 1
		} else {
			if left > 0 {
				d[i] = d[i-1] + 2
				if i-d[i] > 0 {
					d[i] += d[i-d[i]]
				}
				left -= 1
			}
		}
		if d[i] > max {
			max = d[i]
		}
	}
	return max
}
