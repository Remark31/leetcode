package main

import (
	"fmt"
)

import()

func threeSumClosest(nums []int, target int) int {
	
	closest := 999999
	closetarget := 0
	
	// 排序nums,从小到大
	t := nums

	for i := 0 ; i < len(t) ; i ++{
		for j := i + 1 ; j < len(t); j ++{
			if t[i] >= t[j] {
				t[i],t[j] = t[j],t[i]
			}
		}
	}

	for i := 0 ; i < len(t); i ++{
		j := i + 1 
		k := len(t) - 1
		for k > j {
			if t[i] + t[j] + t[k] < target {
				fmt.Println("i , j , k: ", t[i], t[j], t[k])
				close := target - (t[i] + t[j] + t[k])
				if close < closest{
					closest = close
					closetarget = t[i] + t[j] + t[k]
				}
				j += 1
				continue
			}

			if t[i] + t[j] + t[k] > target {
				fmt.Println("i , j , k: ", t[i], t[j], t[k])
				close :=  (t[i] + t[j] + t[k]) - target
				if close < closest{
					closest = close
					closetarget = t[i] + t[j] + t[k]
				}
				k -= 1
				continue
			}

			if t[i] + t[j] + t[k] == target {
				closetarget = target
				return closetarget
			}
			
		}
	}
	
	return closetarget
}

func main(){
	s := []int{-1,2,1,-4}
	t := 1
	fmt.Println(threeSumClosest(s, t))
}