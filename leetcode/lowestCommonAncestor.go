package leetcode

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func initTreeNode(a int) *TreeNode {
	return &TreeNode{Val: a}
}

func makenode(tn *TreeNode, a []int, index int) *TreeNode {
	if index >= len(a) {
		return nil
	}

	if a[index] == -1 {
		return nil
	}
	if tn == nil {
		tn = new(TreeNode)
	}
	tn.Val = a[index]
	fmt.Println(a[index], tn)
	tn.Left = makenode(tn.Left, a, index*2+1)
	tn.Right = makenode(tn.Right, a, index*2+2)

	return tn
}

func InitTreeByArray(a []int) *TreeNode {
	if len(a) == 0 {
		return nil
	}

	var r *TreeNode

	r = makenode(r, a, 0)

	return r
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == p || root == q || root == nil {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	return nil
}
