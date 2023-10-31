package medium

import (
	"fmt"
	"interview_go/datastructure"
)

// https://leetcode.com/problems/lowest-common-ancestor-of-deepest-leaves/
type LowestCommonAncestorOfDeepestLeaves1123 struct {
}

func (l *LowestCommonAncestorOfDeepestLeaves1123) lcaDeepestLeaves(root *datastructure.TreeNode) *datastructure.TreeNode {
	var dfs func(*datastructure.TreeNode) (*datastructure.TreeNode, int)

	dfs = func(node *datastructure.TreeNode) (*datastructure.TreeNode, int) {

		if node == nil {
			return nil, 0
		}

		left, leftDeep := dfs(node.Left)
		right, rightDeep := dfs(node.Right)

		if leftDeep == rightDeep {
			return node, leftDeep + 1
		}

		if leftDeep > rightDeep {
			return left, leftDeep + 1
		}

		return right, rightDeep + 1
	}

	node, _ := l.lcaDeepestLeaveHelper(root)

	return node
}

func (l *LowestCommonAncestorOfDeepestLeaves1123) lcaDeepestLeaveHelper(node *datastructure.TreeNode) (*datastructure.TreeNode, int) {

	if node == nil {
		return nil, 0
	}

	left, leftDepth := l.lcaDeepestLeaveHelper(node.Left)
	right, rightDepth := l.lcaDeepestLeaveHelper(node.Right)

	if leftDepth == rightDepth {
		return node, leftDepth + 1
	}

	if leftDepth > rightDepth {
		return left, leftDepth + 1
	}
	return right, rightDepth + 1
}

func (l *LowestCommonAncestorOfDeepestLeaves1123) Test() {
	left := &datastructure.TreeNode{Val: 5, Left: &datastructure.TreeNode{Val: 6}}
	left.Right = &datastructure.TreeNode{Val: 2, Left: &datastructure.TreeNode{Val: 7}, Right: &datastructure.TreeNode{Val: 4}}
	right := &datastructure.TreeNode{Val: 1, Left: &datastructure.TreeNode{Val: 0}, Right: &datastructure.TreeNode{Val: 8}}
	root := &datastructure.TreeNode{Val: 3, Left: left, Right: right}
	r := l.lcaDeepestLeaves(root)
	e := left.Right

	if r != nil {
		fmt.Printf("r: %d , e: %d -> %v", r.Val, e.Val, r.Val == e.Val)
	} else {
		println("result is nil")
	}
}
