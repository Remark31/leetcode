package main

import()


/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil{
		return head
	}

	before := head
	now := head
	i := 0
	for  ; i < n && now != nil; i ++ {
		now = now.Next
	}

	if i != n {
		return head
	}

	if now == nil {
		head = head.Next
		return head
	}

	for now.Next != nil{
		now = now.Next
		before = before.Next
	}

	before.Next = before.Next.Next

	return head
}