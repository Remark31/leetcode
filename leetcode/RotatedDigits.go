package leetcode

import "strconv"

func isrotednum(num int, changednum, invalidnum []byte) bool {
	s := strconv.Itoa(num)
	haschangednumflag := false
	for i := range s {
		for j := range changednum {
			if s[i] == changednum[j] {
				haschangednumflag = true
			}
		}

		for j := range invalidnum {
			if s[i] == invalidnum[j] {
				return false
			}
		}
	}
	if haschangednumflag {
		return true
	}
	return false
}

func rotatedDigits(N int) int {
	changednum := []byte{'2', '5', '6', '9'}
	invalidnum := []byte{'3', '4', '7'}

	ans := 0
	N = N + 1
	for i := 0; i < N; i++ {
		if isrotednum(i, changednum, invalidnum) {
			ans += 1
		}
	}
	return ans
}
