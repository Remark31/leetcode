package datastruct

import (
	"testing"
)

func initTree(d []int) *TreeNode {
	root := new(TreeNode)
	ir := root
	ir.Val = d[0]
	for i := 1; i < len(d); i++ {
		ir = root
		for {
			if d[i] > ir.Val {
				if ir.Right != nil {
					ir = ir.Right
				} else {
					ir.Right = new(TreeNode)
					ir = ir.Right
					ir.Val = d[i]
					break
				}
			} else {
				if ir.Left != nil {
					ir = ir.Left
				} else {
					ir.Left = new(TreeNode)
					ir = ir.Left
					ir.Val = d[i]
					break
				}
			}
		}
	}

	return root
}

func TestConstructor(t *testing.T) {
	num := []int{3, 4, 6, 7}
	root := initTree(num)
	ir := ConstructorBST(root)
	v := ir.Next()
	if v != 3 {
		t.Error("A", v)
	}
	v = ir.Next()
	if v != 4 {
		t.Error("B", v)
	}
	v = ir.Next()
	if v != 6 {
		t.Error("C", v)
	}
	v = ir.Next()
	if v != 7 {
		t.Error("D", v)
	}

}
