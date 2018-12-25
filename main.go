package main

import (
	"fmt"
	"interview/datastruct"
	"interview/leetcode"
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

func allone() {
	key := []string{"a", "b", "c", "a", "d", "a"}
	obj := datastruct.ConstructorAllOne()
	for _, k := range key {
		obj.Inc(k)
		fmt.Println("Max: ", obj.GetMaxKey())
		fmt.Println("Min: ", obj.GetMinKey())
		obj.Walk()
	}

	for _, k := range key {
		obj.Dec(k)
		fmt.Println("Max: ", obj.GetMaxKey())
		fmt.Println("Min: ", obj.GetMinKey())
		obj.Walk()
	}

}

func minstack() {
	a := []int{3, 5, 4, 2, 7, 8, 3, 1, 1, 4}
	obj := datastruct.ConstructorMinStack()
	for i := range a {
		obj.Push(a[i])
		fmt.Println("i", i, "min: ", obj.GetMin(), "top: ", obj.Top())
	}
	for i := range a {
		fmt.Println("i: ", i, "min: ", obj.GetMin(), "top: ", obj.Top())
		obj.Pop()
	}

}

func lru() {
	obj := datastruct.Constructor(5)
	a := []int{3, 5, 4, 2, 7, 8, 3, 1, 1, 4, 8, 11, 12, 15}
	for i := range a {
		obj.Put(a[i], i)
	}

	for i := 0; i < 5; i++ {
		fmt.Println(obj.Get(i))
	}

}

func lca() {
	a := []int{3, 5, 1, 6, 2, 0, 8, -1, -1, 7, 4}
	r := leetcode.InitTreeByArray(a)

	fmt.Println(r, r.Left, r.Right, r.Left.Left, r.Right.Right)

	l := leetcode.LowestCommonAncestor(r, r.Left, r.Right)

	fmt.Println(l)
}

func main() {
	// heap_sort()
	// queue()
	// fmt.Println(leetcode.ReachNumber(4))
	// allone()
	// minstack()
	// lru()
	lca()
}
