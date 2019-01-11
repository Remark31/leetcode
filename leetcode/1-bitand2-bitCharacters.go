package leetcode

func isOneBitCharacter(bits []int) bool {
	i := 0
	flag := false
	for i < len(bits) {
		if bits[i] == 1 {
			flag = false
			i += 2
		} else if bits[i] == 0 {
			i += 1
			flag = true
		} else {
			return false
		}
	}
	return flag
}
