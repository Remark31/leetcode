package datastruct

import "fmt"

type listnode struct {
	data map[string]bool
	prev *listnode
	next *listnode
}

func (l *listnode) init() {
	l.data = map[string]bool{}
	l.prev = nil
	l.next = nil
}

func (l *listnode) empty() bool {
	return len(l.data) == 0
}

func (l *listnode) getOneKey() string {
	for k, _ := range l.data {
		return k
	}
	return ""
}

func (l *listnode) setKey(k string) {
	l.data[k] = true
}

func (l *listnode) deleteKey(k string) (*listnode, *listnode, *listnode) {
	delete(l.data, k)
	if len(l.data) == 0 {
		if l.prev != nil {
			l.prev.next = l.next
		}
		if l.next != nil {
			l.next.prev = l.prev
		}
	}
	return l.prev, l, l.next
}

func (l *listnode) insertPreNode(k string) *listnode {
	n := new(listnode)
	n.init()
	n.setKey(k)
	if l.prev != nil {
		l.prev.next = n
		n.prev = l.prev
	}
	n.next = l
	l.prev = n
	return n
}

func (l *listnode) insertNextNode(k string) *listnode {
	n := new(listnode)
	n.init()
	n.setKey(k)
	if l.next != nil {
		n.next = l.next
		l.next.prev = n
	}
	n.prev = l
	l.next = n
	return n
}

type mapnode struct {
	value    int
	nodeaddr *listnode
}

type AllOne struct {
	data map[string]*mapnode
	head *listnode
	tail *listnode
}

func (a *AllOne) deletekey(p *listnode, key string) {
	pprev, pnow, pnext := p.deleteKey(key)
	if pnow == a.head {
		if pnow.empty() {
			if pnext != nil {
				a.head = pnext
			}
			if pprev != nil {
				a.head = pprev
			}
		}
	}
	if pnow == a.tail {
		if pnow.empty() {
			if pprev != nil {
				a.tail = pprev
			}
			if pnext != nil {
				a.tail = pnext
			}
		}
	}
}

/** Initialize your data structure here. */
func ConstructorAllOne() AllOne {
	data := map[string]*mapnode{}
	head := new(listnode)
	head.init()
	tail := head
	return AllOne{
		data: data,
		head: head,
		tail: tail,
	}

}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	// 存在，值+1，移动rank的节点位置
	if _, ok := this.data[key]; ok {
		this.data[key].value += 1
		p := this.data[key].nodeaddr

		// 当前排名是否有雷同
		// 前一排名的值
		if p.prev != nil {
			// 前面有人
			pp := p.prev
			k := pp.getOneKey()
			if this.data[key].value == this.data[k].value {
				// 与前一排名相等
				pp.setKey(key)
				this.data[key].nodeaddr = pp
			} else if this.data[key].value < this.data[k].value {
				// 小于前一排名
				n := p.insertPreNode(key)
				this.data[key].nodeaddr = n
			} else {
				// 异常情况
				panic("Inc Over one rank")
			}
			this.deletekey(p, key)
		} else {
			// 已经是头部了
			n := p.insertPreNode(key)
			this.head = n
			this.data[key].nodeaddr = n
			this.deletekey(p, key)
		}
	} else {
		// 不存在，直接从尾部添加
		this.data[key] = new(mapnode)
		this.data[key].value += 1

		if len(this.data) > 1 {
			k := this.tail.getOneKey()
			if this.data[k].value == this.data[key].value {
				this.tail.setKey(key)
				this.data[key].nodeaddr = this.tail
			} else if this.data[k].value > this.data[key].value {
				n := this.tail.insertNextNode(key)
				this.tail = n
				this.data[key].nodeaddr = n
			}
		} else {
			this.tail.setKey(key)
			this.data[key].nodeaddr = this.tail
		}

	}
}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	if _, ok := this.data[key]; ok {
		this.data[key].value -= 1
		if this.data[key].value == 0 {
			// 如果为0，移除
			p := this.data[key].nodeaddr
			this.deletekey(p, key)
			delete(this.data, key)
		} else {
			// 如果不为0，移动排名
			p := this.data[key].nodeaddr
			if p.next != nil {
				pn := p.next
				k := pn.getOneKey()
				if this.data[k].value == this.data[key].value {
					pn.setKey(key)
					this.deletekey(p, key)
					this.data[key].nodeaddr = pn
				} else if this.data[k].value < this.data[key].value {
					n := pn.insertPreNode(key)
					this.deletekey(p, key)
					this.data[key].nodeaddr = n
				} else {
					panic("Dec Over one rank")
				}

			} else {
				n := p.insertNextNode(key)
				this.tail = n
				this.deletekey(p, key)
				this.data[key].nodeaddr = n
			}

		}

	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	return this.head.getOneKey()
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	return this.tail.getOneKey()
}

func (this *AllOne) Walk() {
	for k, v := range this.data {
		fmt.Println("key", k, "value", v.value)
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
