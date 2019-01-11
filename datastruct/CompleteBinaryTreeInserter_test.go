package datastruct

import (
	"fmt"
	"testing"
)

func TestConstructorCBT(t *testing.T) {
	v := 0
	obj := Constructor(nil)
	param_1 := obj.Insert(v)
	param_2 := obj.Get_root()
	fmt.Println("param_1", param_1)
	if param_1 != 0 {
		t.Fatal("Error insert")
	}

	if param_2.Val != 0 || param_2.Left != nil || param_2.Right != nil {
		t.Fatalf("Error Get")
	}

}
