package main

import()

func longestPalindrome(s string) string {
	
	
	max := 0
	maxstring := ""

	// 奇数情况
	for i := 0 ; i < len(s) ;i ++{
		for j := 0 ; ; j ++{
			if i - j < 0 || i + j >= len(s){
				break
			} else {
				if s[i-j] == s[i+j]{
					if len(s[i-j:i+j+1]) >= max{
						max = len(s[i-j:i+j+1])
						maxstring = s[i-j:i+j+1]
					}
				} else{
					break
				}
			}
		}
	}


	// 偶数情况
	for i := 0 ; i < len(s) ; i ++{
		for j := 0 ; ; j ++{
			if i - j < 0 || i + j + 1 >= len(s){
				break
			} else {
				if s[i-j] == s[i+j+1]{
					if len(s[i-j:i+j+2]) >= max{
						max = len(s[i-j:i+j+2])
						maxstring = s[i-j:i+j+2]
					}
				} else{
					break
				}
			}
		}
	}
	return maxstring
}