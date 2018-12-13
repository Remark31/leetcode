package leetcode

import(
	
)

type ListNode struct {
	Val int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil && l2 == nil{
		return nil
	}
	ans := new(ListNode)
	p1 , p2 := l1 , l2
	a := ans
	flag := false

	for p1 != nil {
		if flag {
			a.Val = 1
			flag = false
		}

		if p2 != nil{
			a.Val += p2.Val
			p2 = p2.Next
		}

		a.Val += p1.Val

		if a.Val > 9 {
			a.Val -= 10
			flag = true
		}

		if p1.Next != nil{
			a.Next = new(ListNode)
			a = a.Next
		}
		p1 = p1.Next
	}

	if p2 != nil{
		a.Next = new(ListNode)
		a = a.Next
	}

	for p2 != nil {

		if flag {
			a.Val = 1
			flag = false
		}


		a.Val += p2.Val

		if a.Val > 9 {
			a.Val -= 10
			flag = true
		}

		if p2.Next != nil{
			a.Next = new(ListNode)
			a = a.Next
		}
		p2 = p2.Next
	}

	if flag {
		a.Next = new(ListNode)
		a = a.Next
		a.Val = 1
	}


	return ans
}