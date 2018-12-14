package main

import (
	"fmt"
	"interview/datastruct"
)

func heap_sort() {
	a := []int{4, 5, 6, 2, 3, 4, 8}
	h := datastruct.InitHeap()

	for i := 0; i < len(a); i++ {
		h.Insert(a[i])
	}

	for !h.Empty() {
		fmt.Println("num: ", h.DeleteTop())
	}
}

func queue() {

	a := []int{5, 6, 2, 3, 4, 1}
	q := datastruct.InitQueue()

	for i := 0; i < len(a); i++ {
		q.Push(a[i])
		q.WalkQueue()
	}

	for !q.Empty() {
		fmt.Println(q.Pop())
	}

}

func main() {
	// heap_sort()
	queue()
	// fmt.Println(leetcode.ReachNumber(4))
}
