package medium

import (
	"interview_go/common"
	"interview_go/datastructure"
)

type LowestCommonAncestorOfaBinaryTreeIV struct {
}

// N is the number of nodes in the binary tree
// M is the number of nodes in the nodes array
// H is height of root
// Time complexity : O(N*M)
// Space complexity :O(H)
func (l *LowestCommonAncestorOfaBinaryTreeIV) lowestCommonAncestor(root *datastructure.TreeNode, nodes []*datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil {
		return nil
	}

	for _, node := range nodes {
		if node == root {
			return node
		}
	}
	left := l.lowestCommonAncestor(root.Left, nodes)
	right := l.lowestCommonAncestor(root.Right, nodes)

	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func (l *LowestCommonAncestorOfaBinaryTreeIV) Test() {
	root := &datastructure.TreeNode{Val: 3}
	left := &datastructure.TreeNode{Val: 5}
	left.Left = &datastructure.TreeNode{Val: 6}
	left.Right = &datastructure.TreeNode{Val: 2}
	left.Right.Left = &datastructure.TreeNode{Val: 7}
	left.Right.Right = &datastructure.TreeNode{Val: 4}

	right := &datastructure.TreeNode{Val: 1}
	right.Left = &datastructure.TreeNode{Val: 0}
	right.Right = &datastructure.TreeNode{Val: 8}

	root.Left = left
	root.Right = right

	nodes := []*datastructure.TreeNode{left.Right.Left, left.Right.Right}

	rs := l.lowestCommonAncestor(root, nodes)

	if rs != nil {
		common.PrintInt(rs.Val, left.Right.Val)
	} else {
		println("result is nil")
	}
}
