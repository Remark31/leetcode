package leetcode

func canPartition(nums []int, s map[int]bool, k int, start int, cursum int, curnum int, target int) bool {
	if k == 0 {
		return true
	}
	if cursum == target && curnum > 0 {

		return canPartition(nums, s, k-1, 0, 0, 0, target)
	}
	for i := start; i < len(nums); i++ {
		if s[i] {
			s[i] = false
			if canPartition(nums, s, k, i+1, cursum+nums[i], curnum+1, target) {
				return true
			}
			s[i] = true
		}
	}
	return false
}

func canPartitionKSubsets(nums []int, k int) bool {
	s := make(map[int]bool)
	sum := 0
	for i := range nums {
		sum += nums[i]
		s[i] = true
	}

	average := sum / k
	remain := sum % k
	if remain != 0 {
		return false
	}

	return canPartition(nums, s, k, 0, 0, 0, average)
}
