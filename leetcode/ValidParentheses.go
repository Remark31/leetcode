package leetcode

import()


type stack struct{
	data []byte
}

func (t *stack) push(s byte) {
	t.data = append(t.data, s)
} 

func (t *stack) pop() byte {
	if len(t.data) == 0{
		return '0'
	}

	ans := t.data[len(t.data) - 1]

	t.data = t.data[:len(t.data) - 1]

	return ans
}

func (t *stack) empty() bool{
	return len(t.data) == 0
}

func newStack() *stack {
	return new(stack)
}

func isValid(s string) bool {
	pushop := []byte{'[','(','{'}
	popop := []byte{']',')','}'}

	opmap := map[byte]byte{
		'[':']',
		'(':')',
		'{':'}',
	}

	t := newStack()

	for i := range(s){
		for iop := range(pushop){
			if s[i] == pushop[iop]{
				t.push(s[i])
			}
		}
		for iop := range(popop){
			if s[i] == popop[iop]{
				tmp := t.pop()
				if tmp == '0' {
					return false
				}

				if opmap[tmp] != s[i] {
					return false
				}
			}
		}
	}

	if t.empty(){
		return true
	} else{
		return false
	}



}