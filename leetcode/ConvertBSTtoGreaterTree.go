package leetcode

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 **/

//  [2,0,3,-4,1]

/*
		   2
	    /     \
	   0      3
	  / \
	 -4 1

	    2
	 /     \
	0      3
   / \
  -4 1



*/

// type TreeNode struct {
// 	Val   int
// 	Left  *TreeNode
// 	Right *TreeNode
// }

func makeNode(node *TreeNode, ans *[]int) {

	if node.Right != nil {
		makeNode(node.Right, ans)
	}

	t := node.Val
	for i := 0; i < len(*ans); i++ {
		node.Val += (*ans)[i]
	}

	*ans = append(*ans, t)

	if node.Left != nil {
		makeNode(node.Left, ans)
	}
}

func convertBST(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	ans := []int{}

	makeNode(root, &ans)

	fmt.Println(ans)
	return root
}
