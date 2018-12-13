package leetcode

import()

func Convert(s string, numRows int) string {
	ans := ""
	ansarray := []string{}
	
	for i := 0 ; i < numRows ; i ++{
		ansarray = append(ansarray , "")
	}
	
	is := 0

	for is < len(s) {
		for i := range(ansarray){
			ansarray[i] += s[is:is+1]
			is += 1
			if is == len(s){
				break
			}
		}

		if is == len(s){
			break
		}

		for t := numRows - 2 ; t > 0 ; t --{
			for i := range(ansarray){
				if i == t{
					ansarray[i] += s[is:is+1]
					is += 1
				} else{
					ansarray[i] += " "
				}
				
				if is == len(s){
					break
				}
			}
			
		}
	}


	for _, i := range(ansarray) {
		for j := 0 ; j < len(i) ; j ++{
			if i[j:j+1] != " " {
				ans += i[j:j+1]
			}
		}
	}

	return ans
}