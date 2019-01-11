package leetcode

import (
	"fmt"
	"sort"
)

func lcsmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func lcs(nums []int) int {
	sortnums := make([]int, len(nums))
	copy(sortnums, nums)
	sort.Slice(sortnums, func(i, j int) bool {
		return sortnums[i] < sortnums[j]
	})

	fmt.Println("sort: ", sortnums)

	dp := make([][]int, len(nums)+1)
	for i := 0; i < len(nums)+1; i++ {
		dp[i] = make([]int, len(nums)+1)
	}

	for i := 0; i <= len(nums); i++ {
		for j := 0; j <= len(nums); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else if nums[i-1] == sortnums[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = lcsmax(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	return dp[len(nums)][len(nums)]
}

func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums), 1)
	for i := 0; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = lcsmax(dp[i], dp[j]+1)
			}
		}
	}
	max := 0
	for i := 0; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
