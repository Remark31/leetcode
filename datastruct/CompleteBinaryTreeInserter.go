package datastruct

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type CBTInserter struct {
	data []*TreeNode
}

func (this *CBTInserter) initCBTInserter(node *TreeNode, position int) {
	if node == nil {
		return
	}

	if position >= len(this.data) {
		l := position - len(this.data)
		for i := 0; i <= l; i++ {
			this.data = append(this.data, nil)
		}
	}

	this.data[position] = node
	this.initCBTInserter(node.Left, position*2+1)
	this.initCBTInserter(node.Right, position*2+2)

}

func Constructor(root *TreeNode) CBTInserter {
	ans := CBTInserter{}
	ans.data = []*TreeNode{}
	ans.initCBTInserter(root, 0)
	return ans
}

func (this *CBTInserter) Insert(v int) int {
	p := new(TreeNode)
	p.Val = v
	l := len(this.data)
	if l == 0 {
		this.data = append(this.data, p)
		return v
	}

	parent := this.data[(l-1)/2]
	if parent.Left == nil {
		parent.Left = p
	} else {
		parent.Right = p
	}
	this.data = append(this.data, p)
	return parent.Val
}

func (this *CBTInserter) Get_root() *TreeNode {
	if len(this.data) == 0 {
		return nil
	}
	return this.data[0]
}

/**
 * Your CBTInserter object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Insert(v);
 * param_2 := obj.Get_root();
 */
