package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func walkTree(n *TreeNode, now int) bool {
	if n == nil {
		return true
	} else {
		if now != n.Val {
			return false
		} else {
			return walkTree(n.Left, now) && walkTree(n.Right, now)
		}
	}
}

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return walkTree(root.Left, root.Val) && walkTree(root.Right, root.Val)
}
