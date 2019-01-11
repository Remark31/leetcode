package leetcode

import (
	"strconv"
	"strings"
)

// This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
}

// Return true if this NestedInteger holds a single integer, rather than a nested list.
func (n NestedInteger) IsInteger() bool {
	return true
}

// Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested list
// So before calling this method, you should have a check
func (n NestedInteger) GetInteger() int {
	return 0
}

// Set this NestedInteger to hold a single integer.
func (n *NestedInteger) SetInteger(value int) {}

// Set this NestedInteger to hold a nested list and adds a nested integer to it.
func (n *NestedInteger) Add(elem NestedInteger) {}

// Return the nested list that this NestedInteger holds, if it holds a nested list
// The list length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (n NestedInteger) GetList() []*NestedInteger {
	return nil
}

type neststack struct {
	data []string
}

func (n *neststack) push(d byte) {
	n.data = append(n.data, string(d))
}

func (n *neststack) empty() bool {
	return len(n.data) == 0
}

func (n *neststack) pop() string {
	if n.empty() {
		return ""
	}
	l := len(n.data)
	v := n.data[l-1]
	n.data = n.data[0 : l-1]
	return v
}

func deserializestring(s string, n *NestedInteger) {
	slists := strings.Split(s, ",")
	for _, str := range slists {
		num, _ := strconv.Atoi(str)
		n.SetInteger(num)
	}
}

func deserializesub(s string, begin int, end int, n *NestedInteger) {
	i := begin
	if s[i] != '[' {
		deserializestring(s[begin:end], n)
		return
	}
	i += 1
	ns := new(neststack)
	lastend := i
	newbegin := i
	newend := end
	other := ""
	for i < end {
		if s[i] == '[' {
			if ns.empty() {
				newbegin = i
				other += s[lastend:newbegin]
			}
			ns.push(s[i])
		}
		if s[i] == ']' {
			if !ns.empty() {
				ns.pop()
				if ns.empty() {
					newend = i
					m := new(NestedInteger)
					deserializesub(s, newbegin, newend, m)
					n.Add(*m)
					lastend = i + 1
				}
			}
		}
		i += 1
	}

	deserializestring(other, n)
}

func deserialize(s string) *NestedInteger {
	begin := 0
	end := len(s)
	n := new(NestedInteger)
	deserializesub(s, begin, end, n)

	return n
}
