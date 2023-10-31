package medium

import (
	"fmt"
	"interview_go/datastructure"
)

type LowestCommonAncestorOfaBinaryTreeII_1644 struct {
}

func (l *LowestCommonAncestorOfaBinaryTreeII_1644) lowestCommonAncestor(root, p, q *datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil || p == nil || q == nil {
		return root
	}

	if root == p || root == q {
		return root
	}

	left := l.lowestCommonAncestor(root.Left, p, q)
	right := l.lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

func (s *LowestCommonAncestorOfaBinaryTreeII_1644) Test() {

	left := &datastructure.TreeNode{Val: 5, Left: &datastructure.TreeNode{Val: 6}}
	left.Right = &datastructure.TreeNode{Val: 2, Left: &datastructure.TreeNode{Val: 7}, Right: &datastructure.TreeNode{Val: 4}}
	right := &datastructure.TreeNode{Val: 1, Left: &datastructure.TreeNode{Val: 0}, Right: &datastructure.TreeNode{Val: 8}}
	root := &datastructure.TreeNode{Val: 3, Left: left, Right: right}
	r := s.lowestCommonAncestor(root, left, right)
	e := root

	if r != nil {
		fmt.Printf("r: %d , e: %d -> %v", r.Val, e.Val, r.Val == e.Val)
	} else {
		println("result is nil")
	}
}
