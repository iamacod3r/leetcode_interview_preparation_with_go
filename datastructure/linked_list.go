package datastructure

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (l *LinkedList) AddFront(value int) {
	node := &Node{Value: value}

	if l.head == nil {
		l.head = node
	} else {
		node.Next = l.head
		l.head = node
	}
	l.size++
}

func (l *LinkedList) AddBack(value int) {
	node := &Node{Value: value}

	if l.head == nil {
		l.head = node
	} else {
		curr := l.head
		for curr.Next != nil {
			curr = curr.Next
		}
		curr.Next = node
	}

	l.size++
}

func (l *LinkedList) RemoveFront() error {
	if l.head == nil {
		return fmt.Errorf("remove front error : list is empty")
	}

	l.head = l.head.Next
	l.size--
	return nil
}

func (l *LinkedList) RemoveBack() error {
	if l.head == nil {
		return fmt.Errorf("remove back error: list is empty")
	}

	var prev *Node
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
	l.size--
	return nil
}

func (l *LinkedList) DeleteWithValue(value int) {
	if l.size == 0 || l.head == nil {
		return
	}

	if l.head.Value == value {
		l.head = l.head.Next
		l.size--
		return
	}

	prev := l.head

	for prev.Next.Value != value {
		if prev.Next.Next == nil { // if value not exists
			return
		}
		prev = prev.Next
	}

	prev.Next = prev.Next.Next
	l.size--
}
