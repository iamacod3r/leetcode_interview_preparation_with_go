package easy

import (
	"interview_go/common"
	"interview_go/datastructure"
)

// https://leetcode.com/problems/same-tree/
type SameTree struct {
}

func (s *SameTree) isSameTree(a *datastructure.TreeNode, b *datastructure.TreeNode) bool {

	if a == nil || b == nil {
		return a == nil && b == nil
	}

	return a.Val == b.Val && s.isSameTree(a.Left, b.Left) && s.isSameTree(a.Right, b.Right)
}

func (s *SameTree) check(a, b *datastructure.TreeNode) bool {

	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return a.Val == b.Val
}

func (s *SameTree) isSameTreeIteration(a, b *datastructure.TreeNode) bool {
	if a == nil && b == nil {
		return true
	}
	if !s.check(a, b) {
		return false
	}
	deqA := []*datastructure.TreeNode{}
	deqB := []*datastructure.TreeNode{}
	deqA = append(deqA, a)
	deqB = append(deqB, b)

	for len(deqA) > 0 {
		p := deqA[0]
		q := deqB[0]
		deqA = deqA[1:]
		deqB = deqB[1:]
		if !s.check(p, q) {
			return false
		}
		if !s.check(p.Left, q.Left) {
			return false
		}
		if p.Left != nil {
			deqA = append(deqA, p.Left)
			deqB = append(deqB, q.Left)
		}
		if !s.check(p.Right, q.Right) {
			return false
		}
		if p.Right != nil {
			deqA = append(deqA, p.Right)
			deqB = append(deqB, q.Right)
		}
	}
	return true
}

func (s *SameTree) Test() {
	a := &datastructure.TreeNode{Val: 1, Left: &datastructure.TreeNode{Val: 2}, Right: &datastructure.TreeNode{Val: 3}}
	b := &datastructure.TreeNode{Val: 1, Left: &datastructure.TreeNode{Val: 2}, Right: &datastructure.TreeNode{Val: 3}}

	// b := &datastructure.TreeNode{Val: 1, Left: &datastructure.TreeNode{Val: 2}, Right: &datastructure.TreeNode{Val: 5}}

	e := true
	r := s.isSameTreeIteration(a, b)

	common.PrintBool(r, e)
}
