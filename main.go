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

func main() {
	heap_sort()
}
