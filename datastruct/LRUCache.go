package datastruct

import "fmt"

type LRUnode struct {
	key   int
	value int
	prev  *LRUnode
	next  *LRUnode
}

func (ln *LRUnode) insertNext(k int, v int) *LRUnode {
	n := new(LRUnode)
	n.key = k
	n.value = v

	n.next = ln.next
	ln.next = n

	n.prev = ln
	n.next.prev = n
	return n
}

type LRUCache struct {
	data     map[int]*LRUnode
	capacity int
	head     *LRUnode
	tail     *LRUnode
}

func ConstructorLRUCache(capacity int) LRUCache {
	data := map[int]*LRUnode{}
	cap := capacity
	head := new(LRUnode)
	tail := new(LRUnode)
	head.next = tail
	tail.prev = head

	return LRUCache{data: data, capacity: cap, head: head, tail: tail}
}

func (this *LRUCache) refreash(d *LRUnode) {
	d.prev.next = d.next
	d.next.prev = d.prev

	d.next = this.head.next
	this.head.next = d
	d.prev = this.head

	d.next.prev = d
}

func (this *LRUCache) getoldkey() int {
	n := this.tail.prev
	if n == this.head {
		panic("empty cache")
	}
	return n.key
}

func (this *LRUCache) deletevalue(key int) {
	if d, ok := this.data[key]; ok {
		d.prev.next = d.next
		d.next.prev = d.prev
		delete(this.data, key)
	} else {
		panic("no value to delete")
	}
}

func (this *LRUCache) Get(key int) int {
	if d, ok := this.data[key]; ok {
		this.refreash(d)
		return d.value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.data[key]; ok {
		this.data[key].value = value
		this.refreash(this.data[key])
		return
	}

	if len(this.data) == this.capacity {
		k := this.getoldkey()
		this.deletevalue(k)
		n := this.head.insertNext(key, value)
		this.data[key] = n

	} else {
		n := this.head.insertNext(key, value)
		this.data[key] = n
	}
}

func (this *LRUCache) walkup() {
	for k, v := range this.data {
		fmt.Println("k: ", k, "v", v.value)
	}
	p := this.head.next
	for p != this.tail {
		fmt.Printf("node: %d -> ", p.key)
		p = p.next
	}
	fmt.Println()
	fmt.Println("-----")
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
