package medium

import (
	"fmt"
	"interview_go/datastructure"
)

// https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/description/
type SmallestSubtreeWithAllTheDeepestNodes865 struct {
}

func (s *SmallestSubtreeWithAllTheDeepestNodes865) subtreeWithAllDeepest(root *datastructure.TreeNode) *datastructure.TreeNode {
	var helper func(*datastructure.TreeNode) (*datastructure.TreeNode, int)

	helper = func(node *datastructure.TreeNode) (*datastructure.TreeNode, int) {

		if node == nil {
			return nil, 0
		}

		left, leftDepth := helper(node.Left)
		right, rightDepth := helper(node.Right)

		if leftDepth == rightDepth {
			return node, leftDepth + 1
		}
		if leftDepth > rightDepth {
			return left, leftDepth + 1
		}

		return right, rightDepth + 1
	}

	result, _ := helper(root)

	return result
}

func (s *SmallestSubtreeWithAllTheDeepestNodes865) Test() {
	left := &datastructure.TreeNode{Val: 5, Left: &datastructure.TreeNode{Val: 6}}
	left.Right = &datastructure.TreeNode{Val: 2, Left: &datastructure.TreeNode{Val: 7}, Right: &datastructure.TreeNode{Val: 4}}
	right := &datastructure.TreeNode{Val: 1, Left: &datastructure.TreeNode{Val: 0}, Right: &datastructure.TreeNode{Val: 8}}
	root := &datastructure.TreeNode{Val: 3, Left: left, Right: right}
	r := s.subtreeWithAllDeepest(root)
	e := left.Right

	if r != nil {
		fmt.Printf("r: %d , e: %d -> %v", r.Val, e.Val, r.Val == e.Val)
	} else {
		println("result is nil")
	}
}
