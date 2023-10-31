package easy

import (
	"interview_go/common"
	"interview_go/datastructure"
)

// https://leetcode.com/problems/palindrome-linked-list/
type PalindromeLinkedList234 struct {
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func (p *PalindromeLinkedList234) isPalindrome(head *datastructure.ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// find the middle node
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}

	// reverse the second half
	var prev *datastructure.ListNode
	for slow != nil {
		slow.Next, slow, prev = prev, slow.Next, slow
		// tempNext := slow.Next
		// slow.Next = prev
		// prev = slow
		// slow = tempNext
	}

	// compare the first and second half nodes
	for prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head, prev = head.Next, prev.Next
	}

	return true
}

func (p *PalindromeLinkedList234) Test() {
	head1 := &datastructure.ListNode{Val: 1, Next: &datastructure.ListNode{Val: 2}}
	head1.Next.Next = &datastructure.ListNode{Val: 2, Next: &datastructure.ListNode{Val: 1}}
	e1 := true
	r1 := p.isPalindrome(head1)
	common.PrintBool(r1, e1)

	head2 := &datastructure.ListNode{Val: 1, Next: &datastructure.ListNode{Val: 2}}
	e2 := false
	r2 := p.isPalindrome(head2)
	common.PrintBool(r2, e2)
}
