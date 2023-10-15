package easy

import "interview_go/datastructure"

// https://leetcode.com/problems/middle-of-the-linked-list/description/
type MiddleOfTheLinkedList struct {
}

func (m *MiddleOfTheLinkedList) middleNode(head *datastructure.ListNode) *datastructure.ListNode {

	middle := head
	end := head

	for end != nil && end.Next != nil {
		end = end.Next.Next
		middle = middle.Next
	}

	return middle
}

func (m *MiddleOfTheLinkedList) Test() {
	middle := &datastructure.ListNode{Val: 3, Next: &datastructure.ListNode{Val: 4, Next: &datastructure.ListNode{Val: 5}}}
	head := &datastructure.ListNode{Val: 1, Next: &datastructure.ListNode{Val: 2, Next: middle}}

	r := m.middleNode(head)

	println(r == middle)

	middle = &datastructure.ListNode{Val: 4, Next: &datastructure.ListNode{Val: 5, Next: &datastructure.ListNode{Val: 6}}}
	head = &datastructure.ListNode{Val: 1, Next: &datastructure.ListNode{Val: 2, Next: &datastructure.ListNode{Val: 3, Next: middle}}}

	r = m.middleNode(head)

	println(r == middle)
}
