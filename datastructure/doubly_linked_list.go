package datastructure

import (
	"fmt"
)

type DoubleNode struct {
	Value any
	Prev  *DoubleNode
	Next  *DoubleNode
}

type DoublyLinkedList struct {
	len  int
	head *DoubleNode
	tail *DoubleNode
}

func NewDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (d *DoublyLinkedList) AddHead(value any) {

	newNode := &DoubleNode{
		Value: value,
	}
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		newNode.Next = d.head
		d.head.Prev = newNode
		d.head = newNode
	}

	d.len++
}

func (d *DoublyLinkedList) AddTail(value any) {
	newNode := &DoubleNode{
		Value: value,
	}

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.tail.Next = newNode
		newNode.Prev = d.tail
		d.tail = newNode
	}
}

func (d *DoublyLinkedList) Size() int {
	return d.len
}

func (d *DoublyLinkedList) IsEmpty() bool {
	return d.len == 0
}

func (d *DoublyLinkedList) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("traverseforward error: double linked list is empty")
	}

	curr := d.head
	for curr != nil {
		fmt.Printf("value : %v, prev : %v, next : %v\n", curr.Value, curr.Prev, curr.Next)
		curr = curr.Next
	}
	return nil
}

func (d *DoublyLinkedList) TraverseReverse() error {
	if d.tail == nil {
		return fmt.Errorf("traversereverse error: double linked list is empty")
	}

	curr := d.tail
	for curr != nil {
		fmt.Printf("value : %v, prev : %v, next : %v\n", curr.Value, curr.Prev, curr.Next)
		curr = curr.Prev
	}
	return nil
}

func (d *DoublyLinkedList) PopHead() (any, error) {
	if d.head == nil {
		return nil, fmt.Errorf("pophead error : double linked list is empty")
	}

	item := d.head.Value
	if d.head.Next != nil {
		d.head.Prev = nil
	}

	d.head = d.head.Next
	return item, nil
}

func (d *DoublyLinkedList) PopTail() (any, error) {
	if d.tail == nil {
		return nil, fmt.Errorf("poptail error : double linked list is empty")
	}
	newTail := d.tail.Prev
	item := d.tail.Value
	if newTail != nil {
		newTail.Next = nil
	}
	d.tail.Prev = nil
	d.tail = newTail
	return item, nil
}

func (d *DoublyLinkedList) Delete(value any) error {
	if d.head == nil {
		return fmt.Errorf("delete error : double linked list is empty")
	}

	curr := d.Find(value)
	if curr == nil { // there is no the same value node.
		return nil
	}

	if curr.Prev == nil { // it is head
		_, err := d.PopHead()
		return err
	} else if curr.Next == nil { // it is tail
		_, err := d.PopTail()
		return err
	} else { // it's in the middle point
		newNext := curr.Next
		newPrev := curr.Prev
		newNext.Prev = newPrev
		newPrev.Next = newNext
		curr.Next = nil
		curr.Prev = nil
		curr = nil
	}

	return nil
}

func (d *DoublyLinkedList) Contains(value any) bool {
	curr := d.Find(value)
	return curr != nil
}

func (d *DoublyLinkedList) Find(value any) *DoubleNode {
	curr := d.head
	for curr != nil && curr.Value != value {
		curr = curr.Next
	}

	return curr
}

func (d *DoublyLinkedList) PeekHead() (any, error) {
	if d.head == nil {
		return nil, fmt.Errorf("peekhead error: doubly linked list is empty")
	}
	return d.head.Value, nil
}

func (d *DoublyLinkedList) PeekTail() (any, error) {
	if d.tail == nil {
		return nil, fmt.Errorf("peektail error: doubly linked list is empty")
	}
	return d.tail.Value, nil
}

func (d *DoublyLinkedList) InsertAfter(after any, newValue any) {

	curr := d.head
	for curr != nil && curr.Value != after {
		curr = curr.Next
	}

	if curr == nil || curr.Next == nil {
		d.AddTail(newValue)
	} else {
		newNode := &DoubleNode{
			Value: newValue,
		}

		next := curr.Next
		curr.Next = newNode
		newNode.Prev = curr
		newNode.Next = next
		next.Prev = newNode
	}
}
