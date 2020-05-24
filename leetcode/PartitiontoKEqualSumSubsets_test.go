package leetcode

import (
	"fmt"
	"testing"
)

func TestCanPartitionKSubsets(t *testing.T) {
	s := []int{10, 10, 10, 7, 7, 7, 7, 7, 7, 6, 6, 6}
	n := 3
	fmt.Println("testans: ", canPartitionKSubsets(s, n))
}
