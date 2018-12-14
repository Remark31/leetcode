package leetcode

func removeElement(nums []int, val int) int {
	i := 0
	j := 0

	for j < len(nums) {
		if nums[j] != val {
			nums[i] = nums[j]
			i += 1
			j += 1
		} else {
			for ; j < len(nums) && nums[j] == val; j++ {
			}
		}
	}

	return i
}
