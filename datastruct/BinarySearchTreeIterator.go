package datastruct

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTstackNode struct {
	data int
	next *BSTstackNode
}

type BSTstack struct {
	top *BSTstackNode
}

func (s *BSTstack) push(d int) {
	n := new(BSTstackNode)
	n.data = d
	n.next = s.top
	s.top = n
}

func (s *BSTstack) tops() int {
	if s.top == nil {
		panic("stack is empty")
	}
	v := s.top.data
	s.top = s.top.next
	return v
}

type BSTIterator struct {
	stack *BSTstack
}

func walkTree(node *TreeNode, a *BSTIterator) {
	if node == nil {
		return
	} else {
		walkTree(node.Right, a)
		a.stack.push(node.Val)
		walkTree(node.Left, a)

	}

}

func ConstructorBST(root *TreeNode) BSTIterator {
	head := BSTIterator{}
	head.stack = new(BSTstack)
	walkTree(root, &head)
	return head
}

/** @return the next smallest number */
func (this *BSTIterator) Next() int {
	return this.stack.tops()
}

/** @return whether we have a next smallest number */
func (this *BSTIterator) HasNext() bool {
	return this.stack.top != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
