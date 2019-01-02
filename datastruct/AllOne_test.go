package datastruct

import "testing"

func TestR(t *testing.T) {
	one := ConstructorAllOne()

	one.Inc("aaa")
	one.Inc("bbb")
	one.Inc("ccc")

	one.Inc("aaa")

	v := one.GetMaxKey()
	if v != "aaa" {
		t.Errorf("GetMax Error!")
	}

	one.Dec("bbb")
	v = one.GetMinKey()
	if v != "ccc" {
		t.Errorf("GetMin Error!")
	}
}
