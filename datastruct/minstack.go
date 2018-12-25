package datastruct

type ordernode struct {
	value int
	next  *ordernode
}

func (o *ordernode) getValue() int {
	return o.value
}

func (o *ordernode) insertnext(x int) {
	n := new(ordernode)
	n.value = x
	n.next = o.next
	o.next = n
}

func (o *ordernode) deletenext() {
	if o.next != nil {
		o.next = o.next.next
	}
}

type MinStack struct {
	data []int
	head *ordernode
}

/** initialize your data structure here. */
func ConstructorMinStack() MinStack {
	head := new(ordernode)
	data := []int{}
	return MinStack{data: data, head: head}
}

func (this *MinStack) empty() bool {
	return len(this.data) == 0
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	prev := this.head
	now := this.head.next
	for now != nil {
		if x < now.value {
			prev.insertnext(x)
			return
		}
		prev = now
		now = now.next
	}
	prev.insertnext(x)
}

func (this *MinStack) Pop() {
	l := len(this.data) - 1
	if this.empty() {
		panic("stack is empty")
	}
	v := this.data[l]
	this.data = this.data[:l]

	prev := this.head
	now := this.head.next
	for now != nil {
		if v == now.value {
			prev.deletenext()
			return
		}
		prev = now
		now = now.next
	}
}

func (this *MinStack) Top() int {
	l := len(this.data) - 1
	if this.empty() {
		panic("stack is empty")
	}
	return this.data[l]
}

func (this *MinStack) GetMin() int {
	if this.empty() {
		panic("stack is empty")
	}
	return this.head.next.getValue()
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
