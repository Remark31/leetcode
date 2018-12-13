package main

import()

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == nil {
		return head
	}

	// 换头部第一组
	before := head
	now := head
	after := head.Next

	now.Next = after.Next
	after.Next = now
	head = after

	before = now

	for {
		now = now.Next
		if now == nil {
			break
		}
		after = now.Next
		if after == nil{
			break
		}

		now.Next = after.Next
		after.Next = now
		before.Next = after

		before = now
	}
	


	return head
}