package easy

import "interview_go/datastructure"

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/description/
type RemoveDuplicatesFromSortedList struct {
}

func (r *RemoveDuplicatesFromSortedList) deleteDuplicates(head *datastructure.ListNode) *datastructure.ListNode {
	curr := head
	for curr != nil && curr.Next != nil {
		if curr.Val == curr.Next.Val {
			curr.Next = curr.Next.Next
		} else {
			curr = curr.Next
		}
	}
	return head
}

func (r *RemoveDuplicatesFromSortedList) Test() {
	head := &datastructure.ListNode{}
	head.Val = 1
	head.Next = &datastructure.ListNode{Val: 1}
	head.Next.Next = &datastructure.ListNode{Val: 2}
	result := r.deleteDuplicates(head)

	for result != nil {
		println(result.Val)
		result = result.Next
	}
}
