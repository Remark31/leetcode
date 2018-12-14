package datastruct

import "fmt"

type node struct {
	val  int
	next *node
}

type Queue struct {
	tail  *node
	lenth int
}

func (q *Queue) Empty() bool {
	return q.lenth == 0
}

// []
// [v]
// [v]->[b]

func (q *Queue) Push(v int) {
	if q.Empty() {
		q.tail.val = v
		q.lenth += 1
		return
	}

	t := new(node)
	t.val = v
	t.next = q.tail.next
	q.tail.next = t
	q.tail = q.tail.next
	q.lenth += 1

}

func (q *Queue) Pop() int {
	if q.Empty() {
		panic("Queue is Empty")
	}

	h := q.tail.next
	q.tail.next = q.tail.next.next
	q.lenth -= 1

	return h.val
}

func (q *Queue) WalkQueue() {
	h := q.tail.next
	for i := 0; i < q.lenth; i++ {
		fmt.Printf("i:%d v:%d | ", i, h.val)
		h = h.next
	}
	fmt.Println()
}

func InitQueue() *Queue {
	q := new(node)
	q.next = q

	return &Queue{
		lenth: 0,
		tail:  q,
	}
}
