package leetcode

func numbersof(mid, m, n int) int {
	ans := 0
	for i := 1; i <= m; i++ {
		if mid/i > n {
			ans += n
		} else {
			ans += mid / i
		}
	}
	return ans
}

func findKthNumber(m int, n int, k int) int {
	begin := 1
	end := m * n
	for begin < end {
		mid := (begin + end) / 2
		lens := numbersof(mid, m, n)
		if lens < k {
			begin = mid + 1
		} else {
			end = mid
		}
	}

	return end
}
