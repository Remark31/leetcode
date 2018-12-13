package leetcode

import()

func romanToInt(s string) int {
	c := map[string]int{
		"I" :1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	cc := map[string]int{
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	ans := 0

	i := 0
	for i < len(s) {
		if i + 1 < len(s){
			if t, ok := cc[s[i:i+2]];ok {
				ans += t
				i += 2
				continue	
			} 
		}
		if t , ok := c[s[i:i+1]];ok{
			ans += t
			i+=1
		}
	
	}


	return ans
	
}