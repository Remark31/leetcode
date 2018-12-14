package leetcode

import "fmt"

func RemoveDuplicates(nums []int) int {
	fmt.Println(nums)
	numset := map[int]bool{}
	i := 0
	j := 0

	for j < len(nums) {
		if _, ok := numset[nums[j]]; ok {
			fmt.Println("i", i, "j", j, "nums[j]", nums[j])

			before := nums[j]
			j += 1
			for ; j < len(nums) && before == nums[j]; j++ {
			}

		} else {
			fmt.Println("not repeat", "i", i, "j", j, "nums[j]", nums[j])
			nums[i] = nums[j]
			numset[nums[j]] = true
			i += 1
			j += 1
		}
	}
	fmt.Println(nums)

	return i
}
