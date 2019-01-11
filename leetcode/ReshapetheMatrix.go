package leetcode

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 {
		return nums
	}

	m := len(nums)
	n := len(nums[0])
	if r*c != m*n {
		return nums
	}

	ans := [][]int{}
	for i := 0; i < r; i++ {
		t := make([]int, c)
		ans = append(ans, t)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			ans[i][j] = nums[(i*c+j)/n][(i*c+j)%n]
		}
	}
	return ans
}
