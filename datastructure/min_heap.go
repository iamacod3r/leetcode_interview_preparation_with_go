package datastructure

import "fmt"

type MinHeap struct {
	heap    []int
	size    int
	maxSize int
}

func NewMinHeap(capacity int) *MinHeap {
	return &MinHeap{
		size:    0,
		maxSize: capacity,
		heap:    []int{},
	}
}

func (h *MinHeap) IsLeaf(index int) bool {
	if index >= (h.size/2) && index <= h.size {
		return true
	}
	return false
}

func (h *MinHeap) GetParent(index int) int {
	return (index - 1) / 2
}

func (h *MinHeap) LeftChild(index int) int {
	return 2*index + 1
}

func (h *MinHeap) RightChild(index int) int {
	return 2*index + 2
}

func (h *MinHeap) Insert(item int) error {
	if h.size >= h.maxSize {
		return fmt.Errorf("heap is full")
	}
	h.heap = append(h.heap, item)
	h.size++
	h.upHeapify(h.size - 1)
	return nil
}

func (h *MinHeap) swap(first, second int) {
	h.heap[first], h.heap[second] = h.heap[second], h.heap[first]
}

func (h *MinHeap) upHeapify(index int) {
	parentIndex := h.GetParent(index)
	for h.heap[index] < h.heap[parentIndex] {
		h.swap(index, parentIndex)
		index = parentIndex
	}
}

func (h *MinHeap) downHeapify(current int) {

	if h.IsLeaf(current) {
		return
	}
	smallest := current
	leftChildIndex := h.LeftChild(current)
	rightChildIndex := h.RightChild(current)

	if leftChildIndex < h.size && h.heap[leftChildIndex] < h.heap[smallest] {
		smallest = leftChildIndex
	}
	if rightChildIndex < h.size && h.heap[rightChildIndex] < h.heap[smallest] {
		smallest = rightChildIndex
	}
	if smallest != current {
		h.swap(current, smallest)
		h.downHeapify(smallest)
	}
}

func (h *MinHeap) BuildMinHeap() {
	for index := ((h.size / 2) - 1); index >= 0; index-- {
		h.downHeapify(index)
	}
}

func (h *MinHeap) Remove() int {
	top := h.heap[0]
	h.heap[0] = h.heap[h.size-1]
	h.heap = h.heap[:h.size-1]
	h.size--
	h.downHeapify(0)
	return top
}
