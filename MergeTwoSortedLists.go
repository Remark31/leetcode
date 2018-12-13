package main

import()

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil  {
		return l2
	}
	if l2 == nil {
		return l1
	}

	ans := new(ListNode)
	ta := ans
	
	tl1 := l1
	tl2 := l2

	for tl1 != nil {

		if tl2 != nil && tl1.Val > tl2.Val {
			ta.Val = tl2.Val
			tl2 = tl2.Next
		} else {
			ta.Val = tl1.Val
			tl1 = tl1.Next
		}

		if tl1 != nil{
			ta.Next = new(ListNode)
			ta = ta.Next
		}
	}

	for tl2 != nil {
		ta.Next = new(ListNode)
		ta = ta.Next

		ta.Val = tl2.Val
		tl2 = tl2.Next
	}
	
	return ans
}