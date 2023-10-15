package datastructure

import "fmt"

type MaxHeap struct {
	slice []int
}

func (h *MaxHeap) Insert(key int) {
	h.slice = append(h.slice, key)
	h.maxHeapifyUp(len(h.slice) - 1)
}

func (h *MaxHeap) Extract() int {

	if len(h.slice) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}

	extracted := h.slice[0]
	size := len(h.slice) - 1

	h.slice[0] = h.slice[size]
	h.slice = h.slice[:size]
	h.maxHeapifyDown(0)

	return extracted
}

func (h *MaxHeap) maxHeapifyUp(index int) {

	for h.slice[h.getParent(index)] < h.slice[index] {
		h.swap(h.getParent(index), index)
		index = h.getParent(index)
	}
}

func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.slice) - 1
	l, r := h.getLeft(index), h.getRight(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex || h.slice[l] > h.slice[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.slice[index] < h.slice[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = h.getLeft(index), h.getRight(index)
		} else {
			break
		}
	}
}

func (h *MaxHeap) getParent(index int) int {
	return (index - 1) / 2
}

func (h *MaxHeap) getLeft(index int) int {
	return index*2 + 1
}

func (h *MaxHeap) swap(a, b int) {
	h.slice[a], h.slice[b] = h.slice[b], h.slice[a]
}

func (h *MaxHeap) getRight(index int) int {
	return index*2 + 2
}
