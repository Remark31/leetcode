package datastruct

type allOneListnode struct {
	data  map[string]bool
	value int
	prev  *allOneListnode
	next  *allOneListnode
}

func newallOneListnode() *allOneListnode {
	n := new(allOneListnode)
	n.data = make(map[string]bool)
	return n
}

func (this *allOneListnode) empty() bool {
	return len(this.data) == 0
}

func (this *allOneListnode) insertnext(key string, value int) *allOneListnode {
	if this.next.empty() || this.next.value < value {
		n := newallOneListnode()
		n.value = value
		n.data[key] = true

		n.next = this.next
		this.next = n

		n.prev = this
		n.next.prev = n

	} else {
		n := this.next
		n.data[key] = true
	}

	return this.next
}

func (this *allOneListnode) insertprev(key string, value int) *allOneListnode {
	if this.prev.empty() || this.prev.value > value {
		n := newallOneListnode()

		n.value = value
		n.data[key] = true

		this.prev.next = n
		n.prev = this.prev

		n.next = this
		this.prev = n

	} else {
		n := this.prev
		n.data[key] = true
	}

	return this.prev
}

func (this *allOneListnode) removekey(key string) {
	delete(this.data, key)
	if this.empty() {
		this.prev.next = this.next
		this.next.prev = this.prev
	}
}

type AllOne struct {
	data map[string]*allOneListnode
	head *allOneListnode
	tail *allOneListnode
}

/** Initialize your data structure here. */
func ConstructorAllOne() AllOne {
	a := AllOne{}
	a.data = make(map[string]*allOneListnode)
	a.head = new(allOneListnode)
	a.head.data = make(map[string]bool)

	a.head.next = new(allOneListnode)

	a.tail = a.head.next
	a.tail.prev = a.head
	a.tail.data = make(map[string]bool)

	return a
}

func (this *allOneListnode) getOneKey() string {
	for k, _ := range this.data {
		return k
	}
	return ""
}

func (this *AllOne) empty() bool {
	return this.head.next == this.tail
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	if p, ok := this.data[key]; ok {
		v := p.value + 1
		np := p.insertprev(key, v)
		p.removekey(key)
		this.data[key] = np
	} else {
		p := this.tail.insertprev(key, 1)
		this.data[key] = p
	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if p, ok := this.data[key]; ok {
		v := p.value - 1
		if v != 0 {
			np := p.insertnext(key, v)
			p.removekey(key)
			this.data[key] = np
		} else {
			p.removekey(key)
			delete(this.data, key)
		}

	} else {

	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.empty() {
		return ""
	} else {
		return this.head.next.getOneKey()
	}
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.empty() {
		return ""
	} else {
		return this.tail.prev.getOneKey()
	}
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
