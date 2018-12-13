package main

import()


func Reverse(x int) int {
	flag := true

	MAX := 2147483647
	MIN := -2147483648
	
	if x > MAX || x < MIN{
		return 0
	}


	if x < 0 {
		flag = false
		x = 0-x
	}

	ans := 0
	
	t := x

	for t > 0 {
		ans *= 10
		d := t % 10
		ans += d
		t /= 10
	}
	
	if !flag{
		ans = 0- ans
	}
	
	if ans > MAX || ans < MIN{
		return 0
	}

	return ans
}