package leetcode

func walkzig(n *TreeNode, ans *[][]int, step int) {
	if n == nil {
		return
	}

	if step == len(*ans) {
		*ans = append(*ans, []int{})
	}

	if step%2 == 1 {
		// 正向插入
		(*ans)[step] = append([]int{n.Val}, (*ans)[step]...)
	} else {
		// 背后插入
		(*ans)[step] = append((*ans)[step], n.Val)
	}

	walkzig(n.Left, ans, step+1)
	walkzig(n.Right, ans, step+1)

}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := [][]int{}

	walkzig(root, &ans, 0)

	return ans
}
