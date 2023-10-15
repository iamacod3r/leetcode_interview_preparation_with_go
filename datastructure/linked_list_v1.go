package datastructure

import "fmt"

type Node1 struct {
	Value any
	Next  *Node1
}

type LinkedList1 struct {
	len  int
	head *Node1
}

func NewLinkedList1() *LinkedList1 {
	return &LinkedList1{}
}

func (l *LinkedList1) AddFront(value any) {
	node := &Node1{
		Value: value,
	}

	if l.head == nil {
		l.head = node
	} else {
		node.Next = l.head
		l.head = node
	}
	l.len++
}

func (l *LinkedList1) AddBack(value any) {
	node := &Node1{
		Value: value,
	}

	if l.head == nil {
		l.head = node
	} else {
		curr := l.head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node
	}

	l.len++
}

func (l *LinkedList1) RemoveFront() error {

	if l.head == nil {
		return fmt.Errorf("removefront error : list is empty")
	}

	l.head = l.head.Next
	l.len--
	return nil
}

func (l *LinkedList1) RemoveBack() error {

	if l.head == nil {
		return fmt.Errorf("removefront error : list is empty")
	}

	var prev *Node1
	curr := l.head

	for curr.Next != nil {
		prev = curr
		curr = curr.Next
	}

	if prev != nil {
		prev.Next = nil
	} else {
		l.head = nil
	}

	l.len--
	return nil
}

func (l *LinkedList1) Peek() (any, error) {

	if l.len == 0 || l.head == nil {
		return nil, fmt.Errorf("peek error : linkedlist is empty")
	}
	return l.head.Value, nil
}

func (l *LinkedList1) Size() int {
	return l.len
}

func (l *LinkedList1) Traverse() error {
	if l.head == nil {
		return fmt.Errorf("traverse error: linkedlist is empty")
	}

	curr := l.head
	for curr != nil {
		fmt.Println(curr.Value)
		curr = curr.Next
	}
	return nil
}
