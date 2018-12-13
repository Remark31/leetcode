package leetcode

import(
)


func isRepeatLOS(s string, t string) bool {
	for i := 0 ; i < len(s) ; i ++{
		if t[0] == s[i]{
			return true
		}
	}
	return false
}


func LengthOfLongestSubstring(s string) int {
	max := 0
	
	for i := 0 ; i < len(s) ; i ++ {
		for j := i + 1 ; j <= len(s); j ++{
			if j == len(s) {
				if max < len(s[i:j]){
					max = len(s[i:j])
				}
				break
			}

			if isRepeatLOS(s[i:j], s[j:j+1]){
				if max < len(s[i:j]){
					max = len(s[i:j])
				}
				break	
			}
		}
	}




    return max
}