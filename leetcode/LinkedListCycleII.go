package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	n := head
	nset := make(map[*ListNode]bool)

	for n != nil {
		if _, ok := nset[n]; ok {
			return n
		}
		nset[n] = true
		n = n.Next
	}

	return nil

}
