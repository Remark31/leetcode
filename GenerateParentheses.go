package main

import (
	"fmt"
)

import()


func makeParenthesis(ans *[]string , tans string, bn int, an int) {
	if bn == 0 && an == 0 {
		*ans = append(*ans , tans)
		return
	}

	if bn == 0 {
		s := ""
		for i := 0 ; i < an ; i ++{
			s += ")"
		}
		makeParenthesis(ans , tans + s , bn , 0)	
	}

	if bn == an {
		makeParenthesis(ans , tans + "(" , bn - 1 , an)
	} 

	if bn < an && bn > 0 && an > 0 {
		makeParenthesis(ans , tans + "(" , bn - 1 , an)	
		makeParenthesis(ans , tans + ")", bn , an - 1)	
	}
	
	


} 


func generateParenthesis(n int) []string {
	ans := []string{}

	bn := n 
	an := n


	makeParenthesis(&ans , "", bn , an)



	return ans
}

func main(){
	n := 3
	fmt.Println(generateParenthesis(n))
}
