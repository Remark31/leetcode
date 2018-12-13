package leetcode

import()

const(
	M = 1000
	CM = 900
	D = 500
	CD = 400
	C = 100
	XC = 90
	L = 50
	XL = 40
	X = 10
	IX = 9
	V = 5
	IV = 4
	I = 1
)

func intToRoman(num int) string {
	

	c := map[int]string{
		I:"I" ,
		IV:"IV",
		V:"V",
		IX:"IX",
		X:"X",
		XL:"XL" ,
		L:"L",
		XC:"XC",
		C:"C",
		CD:"CD",
		D:"D",
		CM:"CM",
		M: "M",
	}
	keylist := []int{M,CM,D,CD,C,XC,L,XL,X,IX,V,IV,I}

	a := num
	ans := ""
	i := 0
	for a != 0{
		for i < len(keylist) {
			if a >= keylist[i]{
				ans += c[keylist[i]]
				a -= keylist[i]
			} else {
				i += 1
			}
		}
			
	
	}

	return ans

}