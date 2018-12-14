package datastruct

// 堆的数组实现版本
// 小顶堆

type Heap struct {
	data []int
}

// 插入
func (h *Heap) Insert(d int) {

	h.data = append(h.data, d)

	h.shiftUp(len(h.data) - 1)
}

// 删除
func (h *Heap) DeleteTop() int {
	if h.Empty() {
		panic("Heap is empty, nothing to Delete")
	}

	n := h.data[0]
	end := len(h.data) - 1

	h.data[0], h.data[end] = h.data[end], h.data[0]
	if len(h.data) > 1 {
		h.data = h.data[0:end]
	} else {
		h.data = []int{}
	}

	h.shiftDown(0)

	return n
}

// 判断是否为空
func (h *Heap) Empty() bool {
	if len(h.data) == 0 {
		return true
	}
	return false
}

// 堆调整
func (h *Heap) reheap() {

}

// 上浮
func (h *Heap) shiftUp(index int) {
	parent := (index - 1) / 2

	for index > 0 {
		if h.data[index] < h.data[parent] {
			h.data[index], h.data[parent] = h.data[parent], h.data[index]
		}
		index = parent
		parent = (index - 1) / 2
	}

}

// 下沉
func (h *Heap) shiftDown(index int) {
	lchild := 2*index + 1
	rchild := 2*index + 2
	n := len(h.data)
	for lchild < n {
		swapindex := lchild
		if rchild < n && h.data[lchild] > h.data[rchild] {
			swapindex = rchild
		}

		h.data[index], h.data[swapindex] = h.data[swapindex], h.data[index]

		index = swapindex
		lchild = 2*index + 1
		rchild = 2*index + 2
	}
}

// 初始化
func InitHeap() *Heap {
	return new(Heap)
}
