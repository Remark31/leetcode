package leetcode

// DFS

/*
func alltrue(v map[int]bool) bool {
	ans := true
	for _, value := range v {
		ans = ans && value
	}
	return ans
}

func shunzi(nums []int, visited map[int]bool, start int, now int, curnum int) bool {
	// 退出条件
	if alltrue(visited) && curnum >= 2 {
		return true
	}

	for i := start; i < len(nums); i++ {
		if !visited[i] {
			// 未被访问过
			if now != 0 {
				for i < len(nums) && nums[i] != now+1 {
					i += 1
				}
				if i < len(nums) && nums[i] == now+1 {
					visited[i] = true
					if shunzi(nums, visited, i+1, nums[i], curnum+1) {
						return true
					}
					visited[i] = false
				} else if curnum >= 2 {
					return shunzi(nums, visited, 0, 0, 0)
				} else {
					continue
				}
			}
			if now == 0 {
				visited[i] = true
				if shunzi(nums, visited, i+1, nums[i], curnum+1) {
					return true
				}
				visited[i] = false
			}
		}
	}

	return false
}

func isPossible(nums []int) bool {
	visited := make(map[int]bool)

	for i := range nums {
		visited[i] = false
	}

	return shunzi(nums, visited, 0, 0, 0)

}
*/

func isPossible(nums []int) bool {

}
