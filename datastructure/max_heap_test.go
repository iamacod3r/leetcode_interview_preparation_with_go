package datastructure

import (
	"fmt"
	"testing"
)

func TestMaxHeap(t *testing.T) {

	m := &MaxHeap{}

	buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}

	for _, v := range buildHeap {
		m.Insert(v)
	}
	fmt.Println(m)

	for i := 0; i < 5; i++ {
		m.Extract()
		fmt.Println(m)
	}
}
