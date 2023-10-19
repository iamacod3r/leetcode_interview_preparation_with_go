package medium

import (
	"interview_go/common"
	"interview_go/datastructure"
)

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree/editorial/
type LowestCommonAncestorOfaBinaryTree236 struct{}

type NodePair struct {
	Node        *datastructure.TreeNode
	ParentState int
}

// N is the number of nodes in the binary tree
// Time Complexity : O(n)
// Space Complexity : O(n)
func (w *LowestCommonAncestorOfaBinaryTree236) lowestCommonAncestorIterativeWithoutParents(root, p, q *datastructure.TreeNode) *datastructure.TreeNode {
	bothPending, bothDone := 2, 0
	oneFound := false
	stack := []NodePair{{root, bothPending}}

	var lca *datastructure.TreeNode
	var child *datastructure.TreeNode

	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if top.ParentState != bothDone {
			if top.ParentState == bothPending {

				if top.Node == p || top.Node == q {
					if oneFound {
						return lca
					} else {
						oneFound = true
						lca = top.Node
					}
				}
				child = top.Node.Left
			} else {
				child = top.Node.Right
			}

			stack = append(stack, NodePair{top.Node, top.ParentState - 1})
			if child != nil {
				stack = append(stack, NodePair{child, bothPending})
			}
		} else {
			if lca == top.Node && oneFound {
				lca = stack[len(stack)-1].Node
			}
		}
	}
	return nil
}

// N is the number of nodes in the binary tree
// Time Complexity : O(n)
// Space Complexity : O(n)
func (w *LowestCommonAncestorOfaBinaryTree236) lowestCommonAncestorIterativeWithParents(root, p, q *datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil {
		return nil
	}

	nodeStack := []*datastructure.TreeNode{root}
	parentMap := map[*datastructure.TreeNode]*datastructure.TreeNode{root: nil}
	for len(nodeStack) > 0 {
		node := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		if node.Left != nil {
			parentMap[node.Left] = node
			nodeStack = append(nodeStack, node.Left)
		}

		if node.Right != nil {
			parentMap[node.Right] = node
			nodeStack = append(nodeStack, node.Right)
		}
	}

	ancestors := map[*datastructure.TreeNode]bool{}
	for p != nil {
		ancestors[p] = true
		p = parentMap[p]
	}

	for q != nil {
		if ancestors[q] {
			return q
		}
		q = parentMap[q]
	}

	return nil
}

// N is the number of nodes in the binary tree
// Time Complexity : O(n)
// Space Complexity : O(n)
func (w *LowestCommonAncestorOfaBinaryTree236) lowestCommonAncestorRecursive(root, p, q *datastructure.TreeNode) *datastructure.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}

	left := w.lowestCommonAncestorRecursive(root.Left, p, q)
	right := w.lowestCommonAncestorRecursive(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}

func (w *LowestCommonAncestorOfaBinaryTree236) Test() {

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
	rs := w.lowestCommonAncestorRecursive(root, left.Right.Right, right.Left)
	e := 3
	if rs != nil {
		common.PrintInt(rs.Val, e)
	} else {
		println("result is nil")
	}

}
