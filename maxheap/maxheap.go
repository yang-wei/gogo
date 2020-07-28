package maxheap

type maxHeap struct {
	arr []int
}

func NewMaxHeap(arr []int) *maxHeap {
	mh := &maxHeap{
		arr: arr,
	}
	mh.buildMaxHeap()
	return mh
}

// i is 1-based i
func (h *maxHeap) heapify(i int) {

}

func (h *maxHeap) buildMaxHeap() {
	for i := len(h.arr) / 2; i > 0; i-- {
		h.heapify(i)
	}
}

func (h *maxHeap) insert(node int) {
	h.arr = append(h.arr, node)
	insertedPosition := len(h.arr)
	for i := insertedPosition; i > 1 && h.arr[i-1] > h.arr[i/2-1]; i = i / 2 {
		h.arr[i-1], h.arr[i/2-1] = h.arr[i/2-1], h.arr[i-1]
	}
}

func (h *maxHeap) extractMax() int {

	max := h.getMax()
	if len(h.arr) > 0 {
		tail := len(h.arr) - 1
		h.arr[0], h.arr[tail] = h.arr[tail], h.arr[0]
		h.arr = h.arr[:tail]
		h.heapify(1)
	}
	return max

}

func (h *maxHeap) getMax() int {
	if len(h.arr) > 0 {
		return h.arr[0]
	}
	return -1
}

func (h *maxHeap) sort() {
	len := len(h.arr)
	result := make([]int, len, len)
	for i := 0; i < len; i++ {
		result[i] = h.extractMax()
	}
	h.arr = result
}
