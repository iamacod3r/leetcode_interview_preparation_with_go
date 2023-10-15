package datastructure

import "fmt"

type MaxHeap1 struct {
	heap    []int
	size    int
	maxSize int
}

func NewMaxHeap1(capacity int) *MaxHeap1 {
	return &MaxHeap1{
		heap:    []int{},
		size:    0,
		maxSize: capacity,
	}
}

func (h *MaxHeap1) IsLeaf(index int) bool {
	if index >= (h.size/2) && index <= h.size {
		return true
	}
	return false
}

func (h *MaxHeap1) GetParent(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap1) LeftChild(index int) int {
	return 2*index + 1
}

func (h *MaxHeap1) RightChild(index int) int {
	return 2*index + 2
}

func (h *MaxHeap1) Insert(item int) error {
	if h.size > h.maxSize {
		return fmt.Errorf("heap is full")
	}

	h.heap = append(h.heap, item)
	h.size++
	h.upHeapify(h.size - 1)
	return nil
}

func (h *MaxHeap1) swap(first, second int) {
	h.heap[first], h.heap[second] = h.heap[second], h.heap[first]
}

func (h *MaxHeap1) upHeapify(index int) {
	parent := h.GetParent(index)
	for h.heap[index] > h.heap[parent] {
		h.swap(index, parent)
		index = parent
	}
}

func (h *MaxHeap1) downHeapify(index int) {
	if h.IsLeaf(index) {
		return
	}

	largest := index
	leftChildIndex := h.LeftChild(index)
	rightChildIndex := h.RightChild(index)

	if leftChildIndex < h.size && h.heap[leftChildIndex] > h.heap[largest] {
		largest = leftChildIndex
	}

	if rightChildIndex < h.size && h.heap[rightChildIndex] > h.heap[largest] {
		largest = rightChildIndex
	}

	if largest != index {
		h.swap(index, largest)
		h.downHeapify(largest)
	}
}

func (h *MaxHeap1) BuildMaxHeap() {
	for index := ((h.size / 2) - 1); index >= 0; index-- {
		h.downHeapify(index)
	}
}

func (h *MaxHeap1) Remove() int {
	top := h.heap[0]
	h.heap[0] = h.heap[h.size-1]
	h.heap = h.heap[:h.size-1]
	h.size--
	h.downHeapify(0)
	return top
}
