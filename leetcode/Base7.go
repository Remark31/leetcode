package leetcode

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}
	ans := ""
	rans := ""
	flag := true
	if num < 0 {
		num = 0 - num
		flag = false
	}

	for num != 0 {
		t := num % 7
		ans += string('0' + t)
		num = num / 7
	}

	if !flag {
		ans += "-"
	}

	for i := len(ans) - 1; i >= 0; i-- {
		rans += string(ans[i])
	}

	return rans
}
