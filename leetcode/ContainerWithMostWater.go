package leetcode

import()


func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	max := 0

	for i := 0 ;i < len(height) ; i ++{
		for j := len(height) - 1; j > i ; j --{
			area := min(height[i],height[j])*(j-i)
			if area > max{
				max = area
			}
		}
	}
	
	return max
}

