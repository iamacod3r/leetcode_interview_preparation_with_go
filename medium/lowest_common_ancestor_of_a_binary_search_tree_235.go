package medium

import (
	"interview_go/common"
	"interview_go/datastructure"
)

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/description/
type LowestCommonAncestorOfaBinarySearchTree235 struct {
}

// Time complexity : O(n)
// Space complexity : O(1)
func (t *LowestCommonAncestorOfaBinarySearchTree235) lowestCommonAncestorIterative(root, p, q *datastructure.TreeNode) *datastructure.TreeNode {

	if root == nil || p == nil || q == nil || root == q || root == p {
		return root
	}

	node := root
	for node != nil {

		if p.Val > node.Val && q.Val > node.Val {
			node = node.Right
		} else if p.Val < node.Val && q.Val < node.Val {
			node = node.Left
		} else {
			return node
		}

	}
	return nil
}

// Time complexity : O(n)
// Space complexity : O(n)
func (t *LowestCommonAncestorOfaBinarySearchTree235) lowestCommonAncestorRecursive(root, p, q *datastructure.TreeNode) *datastructure.TreeNode {

	if root == nil || p == nil || q == nil || root == q || root == p {
		return root
	}

	if p.Val < root.Val && q.Val < root.Val {
		return t.lowestCommonAncestorRecursive(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return t.lowestCommonAncestorRecursive(root.Right, p, q)
	}
	return root
}

func (t *LowestCommonAncestorOfaBinarySearchTree235) Test() {

	root := &datastructure.TreeNode{Val: 6}
	left := &datastructure.TreeNode{Val: 2}
	left.Left = &datastructure.TreeNode{Val: 0}
	left.Right = &datastructure.TreeNode{Val: 4}
	left.Right.Left = &datastructure.TreeNode{Val: 3}
	left.Right.Right = &datastructure.TreeNode{Val: 5}

	right := &datastructure.TreeNode{Val: 8}
	right.Left = &datastructure.TreeNode{Val: 7}
	right.Right = &datastructure.TreeNode{Val: 9}

	root.Left = left
	root.Right = right
	rs := t.lowestCommonAncestorIterative(root, left.Right.Right, left.Right)
	e := 4
	if rs != nil {
		common.PrintInt(rs.Val, e)
	} else {
		println("result is nil")
	}

}
