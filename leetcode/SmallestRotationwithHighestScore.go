package leetcode

func getscore(a []int) int {
	ans := 0
	for i := range a {
		if a[i] <= i {
			ans += 1
		}
	}
	return ans
}

func bestRotation(A []int) int {
	max := 0
	maxindex := 0

	for i := 0; i < len(A); i++ {
		t := append(A[i:len(A)], A[0:i]...)
		s := getscore(t)
		if s > max {
			maxindex = i
			max = s
		}
	}
	return maxindex
}
