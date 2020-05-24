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
	unused := make(map[int]int)
	need := make(map[int]int)
	for i := range nums{
		if _ , ok := unused[nums[i]]; !ok{
			unused[nums[i]] = 0
		}
		unused[nums[i]] += 1
	}

	for i := range nums{
        v , ok := need[nums[i]]
		v1 , ok1 := unused[nums[i]+1]
		v2 , ok2 := unused[nums[i]+2]
		if unused[nums[i]] == 0{
			continue
		} else if ok && v > 0{
			need[nums[i]] -= 1
			if _ , k := need[nums[i]+1]; !k {
				need[nums[i]+1] = 0
			}
			need[nums[i]+1] += 1
			unused[nums[i]] -= 1
		} else if ok1 && v1 > 0 && ok2 && v2 > 0{
			if _ , ok3 := need[nums[i]+3] ; !ok3{
				need[nums[i]+3] = 0
			} 
			need[nums[i]+3] += 1
            unused[nums[i]] -= 1
			unused[nums[i]+1] -= 1
			unused[nums[i]+2] -= 1
		} else{
			return false
		}

	}
	return true
}
