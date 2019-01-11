package leetcode

import (
	"fmt"
	"testing"
)

func TestLcs(t *testing.T) {
	a := []int{10, 9, 2, 5, 3, 7, 101, 18}
	v := lengthOfLIS(a)
	if v != 4 {
		fmt.Println("Error ", v)
		t.Fatal("BUG")
	}
}
