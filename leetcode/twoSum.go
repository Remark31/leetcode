package leetcode

import(
)


func TwoSum(nums []int, target int) []int {
	i , j := 0, 0
	for i = 0; i < len(nums); i++  {
		for j = i + 1 ; j < len(nums); j ++ {
			if nums[i] + nums[j] == target{
				return []int{i , j}
			}
		}
	}
	return []int{}
}

