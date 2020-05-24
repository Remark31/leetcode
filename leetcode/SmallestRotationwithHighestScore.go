package leetcode

<<<<<<< HEAD
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
=======
func bestRotation(A []int) int {
	return 0
>>>>>>> 6b4d0b46e6b47ba4c2229521978a743f602d61f9
}
