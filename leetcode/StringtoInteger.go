package leetcode

import()

func isInterger(a byte) bool{
	if a > 47 && a < 58 {
		return true
	}
	return false
}



func myAtoi(str string) int {
	
	MAX := 2147483647

	flagJudge := false
	flag := true
	ans := 0

	for i := range str{
		
		// 空格
		if !flagJudge && str[i] == ' ' {
			continue
		}
		

		// 符号
		if !flagJudge {
			if str[i] == '+' {
				flag = true
				flagJudge = true
				continue
			} else if str[i] == '-' {
				flag = false
				flagJudge = true
				continue
			}
		}



		// 数字
		if isInterger(str[i]){

			ans *=10
			ans += int(str[i] - 48)
			
			flagJudge = true


			if ans > MAX && flag {
				ans = MAX
				break
			}

			if ans > MAX + 1 && !flag{
				ans = MAX + 1
				break
			}
			
			
		} else{
			break
		}



	}

	

	if !flag {
		ans =  0 - ans
	}

	return ans

}