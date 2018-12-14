package leetcode

func isRepeat3Sum(anss [][]int, ans []int) bool {

	for s := range anss {
		flag := false
		for i := 0; i < len(ans); i++ {
			if anss[s][i] != ans[i] {
				flag = true
				break
			}
		}
		if !flag {
			return true
		}
	}
	return false
}

func fourSum(nums []int, target int) [][]int {
	ans := [][]int{}

	t := nums

	for i := 0; i < len(t); i++ {
		for j := i + 1; j < len(t); j++ {
			if t[i] > t[j] {
				t[i], t[j] = t[j], t[i]
			}
		}
	}

	for i := 0; i < len(t); i++ {
		for j := i + 1; j < len(t); j++ {
			k := j + 1
			l := len(t) - 1

			for k < l {
				if t[i]+t[j]+t[k]+t[l] < target {
					k += 1
					continue
				}
				if t[i]+t[j]+t[k]+t[l] > target {
					l -= 1
					continue
				}

				if t[i]+t[j]+t[k]+t[l] == target {
					if !isRepeat3Sum(ans, []int{t[i], t[j], t[k], t[l]}) {
						ans = append(ans, []int{t[i], t[j], t[k], t[l]})
					}

					for k < l && t[k] == t[k+1] {
						k += 1
					}
					k += 1
				}

			}

		}
	}

	return ans

}
