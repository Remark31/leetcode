package leetcode

import "strconv"

const (
	ADD   = "+"
	MINUS = "-"
	MULI  = "*"
	DIV   = "/"
)

type tstack struct {
	data []int
}

func newtStack() *tstack {
	return new(tstack)
}

func (t *tstack) push(a int) {
	t.data = append(t.data, a)
}

func (t *tstack) empty() bool {
	return len(t.data) == 0
}

func (t *tstack) pop() int {
	if t.empty() {
		panic("stack is empty")
	}

	l := len(t.data)

	d := t.data[l-1]

	t.data = t.data[:l-1]

	return d

}

func evalRPN(tokens []string) int {
	t := newtStack()

	for i := range tokens {
		switch tokens[i] {
		case ADD:
			a := t.pop()
			b := t.pop()
			t.push(b + a)

		case MINUS:
			a := t.pop()
			b := t.pop()
			t.push(b - a)

		case MULI:
			a := t.pop()
			b := t.pop()
			t.push(b * a)

		case DIV:

			a := t.pop()
			b := t.pop()
			t.push(b / a)

		default:
			n, err := strconv.Atoi(tokens[i])
			if err != nil {
				panic("input is illega")
			}
			t.push(n)
		}
	}
	return t.pop()

}
