package main
import()

func combinenumber(a []string , b []string) []string {
	ans := []string{}
	
	if len(a) == 0{
		return b
	}

	for ia := range(a){
		for ib := range(b){
			t := a[ia] + b[ib]
			ans = append(ans, t)
		}
	}
	return ans
}

func letterCombinations(digits string) []string {
	numbermap := map[string][]string{}
	numbermap["2"] = []string{"a","b","c"}
	numbermap["3"] = []string{"d","e","f"}
	numbermap["4"] = []string{"g","h","i"}
	numbermap["5"] = []string{"j","k","l"}
	numbermap["6"] = []string{"m","n","o"}
	numbermap["7"] = []string{"p","q","r","s"}
	numbermap["8"] = []string{"t","u","v"}
	numbermap["9"] = []string{"w","x","y","z"}

	ans := []string{}

	for i := range(digits){
		ans = combinenumber(ans , numbermap[digits[i:i+1]])	
	}

	return ans
	
}