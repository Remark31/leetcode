package leetcode

type equip struct {
	odd  map[byte]int
	even map[byte]int
}

func initequip(a string) *equip {
	e := new(equip)
	e.even = make(map[byte]int)
	e.odd = make(map[byte]int)

	for i := range a {
		if i%2 == 0 {
			if _, ok := e.even[a[i]]; !ok {
				e.even[a[i]] = 1
			} else {
				e.even[a[i]] += 1
			}
		} else {
			if _, ok := e.odd[a[i]]; !ok {
				e.odd[a[i]] = 1
			} else {
				e.odd[a[i]] += 1
			}
		}
	}
	return e
}

func compareequip(a *equip, b *equip) bool {
	for k, v := range a.even {
		if _, ok := b.even[k]; ok {
			if b.even[k] != v {
				return false
			}
		} else {
			return false
		}
	}

	for k, v := range a.odd {
		if _, ok := b.odd[k]; ok {
			if b.odd[k] != v {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

func numSpecialEquivGroups(A []string) int {
	es := []*equip{}

	for i := range A {
		e := initequip(A[i])
		es = append(es, e)
	}

	num := make([]int, len(A))
	ans := 0

	for i := 0; i < len(A); i++ {
		if num[i] == 0 {
			num[i] = 1
			ans += 1
			for j := 1; j < len(A); j++ {
				if num[j] == 0 && compareequip(es[i], es[j]) {
					num[j] = 1
				}
			}
		}
	}
	return ans
}
