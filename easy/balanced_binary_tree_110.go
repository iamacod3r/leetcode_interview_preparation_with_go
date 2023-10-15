package easy

import (
	"interview_go/common"
	"interview_go/datastructure"
	"math"
)

// https://leetcode.com/problems/balanced-binary-tree/

/*
if recursive calls BEFORE conditional check, then its BOTTOM UP.
If recursive call AFTER conditional check, its TOP DOWN.
*/
type BalancedBinaryTree struct {
}

// best solution
func (b *BalancedBinaryTree) getHeight(node *datastructure.TreeNode) int {
	if node == nil {
		return 0
	}

	left := b.getHeight(node.Left)
	if left == -1 {
		return -1
	}
	right := b.getHeight(node.Right)
	if right == -1 || math.Abs(float64(left-right)) > 1 {
		return -1
	}

	return int(math.Max(float64(left), float64(right))) + 1
}

// best solution
func (b *BalancedBinaryTree) isBalanced(root *datastructure.TreeNode) bool {
	return b.getHeight(root) != -1
}

type TreeInfo struct {
	height   int
	balanced bool
}

func (b *BalancedBinaryTree) isBalancedTreeHelper(root *datastructure.TreeNode) TreeInfo {
	if root == nil {
		return TreeInfo{-1, true}
	}

	left := b.isBalancedTreeHelper(root.Left)
	if !left.balanced {
		return TreeInfo{-1, false}
	}
	right := b.isBalancedTreeHelper(root.Right)
	if !right.balanced {
		return TreeInfo{-1, false}
	}

	if math.Abs(float64(left.height-right.height)) < 2 {
		return TreeInfo{int(math.Max(float64(left.height), float64(right.height))) + 1, true}
	}

	return TreeInfo{-1, false}
}

func (b *BalancedBinaryTree) isBalancedWithHelper(root *datastructure.TreeNode) bool {
	return b.isBalancedTreeHelper(root).balanced
}

func (b *BalancedBinaryTree) height_nLogN(root *datastructure.TreeNode) int {
	if root == nil {
		return -1
	}
	return 1 + int(math.Max(float64(b.height_nLogN(root.Left)), float64(b.height_nLogN(root.Right))))
}

// time complexity : O(nlogn)
// space complexity : O(n)
func (b *BalancedBinaryTree) isBalanced_nLogN(root *datastructure.TreeNode) bool {
	if root == nil {
		return true
	}

	leftheight := b.height_nLogN(root.Left)
	rightHeight := b.height_nLogN(root.Right)

	leftBalanced := b.isBalanced_nLogN(root.Left)
	rightBalanced := b.isBalanced_nLogN(root.Right)

	return math.Abs(float64(leftheight-rightHeight)) < 2 && leftBalanced && rightBalanced
}

func (b *BalancedBinaryTree) Test() {

	// [3,9,20,null,null,15,7]
	root := &datastructure.TreeNode{Val: 3}
	root.Left = &datastructure.TreeNode{Val: 9}
	rootRight := &datastructure.TreeNode{Val: 20}
	rootRight.Left = &datastructure.TreeNode{Val: 15}
	rootRight.Right = &datastructure.TreeNode{Val: 7}
	root.Right = rootRight
	e := true
	r := b.isBalanced(root)
	common.PrintBool(r, e)

	// [1,2,2,3,3,null,null,4,4]
	next := &datastructure.TreeNode{Val: 1, Right: &datastructure.TreeNode{Val: 2}}
	mostLeft := &datastructure.TreeNode{Val: 3, Left: &datastructure.TreeNode{Val: 4}, Right: &datastructure.TreeNode{Val: 4}}
	next.Left = &datastructure.TreeNode{Val: 2, Left: mostLeft, Right: &datastructure.TreeNode{Val: 3}}
	ne := false
	re := b.isBalanced(next)
	common.PrintBool(re, ne)
}
