package leetcode

import (
	"fmt"
	"testing"
)

func TestFindReplaceString(t *testing.T) {

	S := "jjievdtjfb"
	indexs := []int{4, 6, 1}
	sources := []string{"md", "tjgb", "jf"}
	targets := []string{"foe", "oov", "e"}

	fmt.Println(findReplaceString(S, indexs, sources, targets))
}
