package main

import("fmt")

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	
	ans := ""
	
	i := 0
	flag := false

	for {
		if i == len(strs[0]){
			break
		}
		before := strs[0][i]
		for _, s := range(strs){
			if i == len(s){
				flag = true
				break
			}
			
			if before != s[i]{
				flag = true
				break
			}
		}


		if flag{
			break
		}
		ans += string(before)
		i += 1
	}

	return ans
}

func main() {
	str := []string{"cba",""}
	fmt.Println(longestCommonPrefix(str))
}