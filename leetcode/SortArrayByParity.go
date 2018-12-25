package leetcode

func sortArrayByParity(A []int) []int {
	lenth := len(A)
	if lenth == 0 {
		return A
	}
	i, j := 0, lenth-1
	// 偶数在前奇数在后
	for i < j {
		for i < lenth && A[i]%2 == 0 {
			i += 1
		}
		for j > 0 && A[j]%2 != 0 {
			j -= 1
		}
		if i < j {
			A[i], A[j] = A[j], A[i]
			i += 1
			j -= 1
		}
	}
	return A
}
