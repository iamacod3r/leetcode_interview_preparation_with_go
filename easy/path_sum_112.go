package easy

import (
	"interview_go/common"
	"interview_go/datastructure"
)

// https://leetcode.com/problems/path-sum/
type PathSum struct{}

// In both methods
// Time Complexities : O(N)
// Space Complexities : best (balanced Tree) -> O(logN) | worst (unbalanced Tree) ->O(N)

func (p *PathSum) hasPathSumRecursion(root *datastructure.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	targetSum -= root.Val

	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}

	return p.hasPathSumRecursion(root.Left, targetSum) || p.hasPathSumRecursion(root.Right, targetSum)
}

// Iteration
func (p *PathSum) hasPathSumDFS(root *datastructure.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	stack := []*datastructure.TreeNode{root}
	sumStack := []int{targetSum - root.Val}

	for len(stack) > 0 {
		// get last Node from stack
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// get last currSum from sumStack
		currSum := sumStack[len(sumStack)-1]
		sumStack = sumStack[:len(sumStack)-1]

		if node.Left == nil && node.Right == nil && currSum == 0 {
			return true
		}

		if node.Left != nil {
			stack = append(stack, node.Left)
			sumStack = append(sumStack, currSum-node.Left.Val)
		}

		if node.Right != nil {
			stack = append(stack, node.Right)
			sumStack = append(sumStack, currSum-node.Right.Val)
		}
	}
	return false
}

func (p *PathSum) Test() {
	// root := &datastructure.TreeNode{Val: 5}
	// root.Left = &datastructure.TreeNode{Val: 4}
	// rootLeft := &datastructure.TreeNode{Val: 11}
	// rootLeft.Left = &datastructure.TreeNode{Val: 7}
	// rootLeft.Right = &datastructure.TreeNode{Val: 2}
	// root.Left.Left = rootLeft
	// rootRight := &datastructure.TreeNode{Val: 8}
	// rootRight.Left = &datastructure.TreeNode{Val: 13}
	// rootRight.Right = &datastructure.TreeNode{Val: 4}
	// rootRight.Right.Right = &datastructure.TreeNode{Val: 1}
	// root.Right = rootRight

	// e := true
	// r := p.hasPathSumDFS(root, 22)
	// common.PrintBool(r, e)

	// root2 := &datastructure.TreeNode{Val: 1}
	// root2.Left = &datastructure.TreeNode{Val: 2}
	// root2.Right = &datastructure.TreeNode{Val: 3}
	// e2 := false
	// r2 := p.hasPathSumDFS(root2, 5)
	// common.PrintBool(r2, e2)

	root3 := &datastructure.TreeNode{Val: 1}
	e3 := true
	r3 := p.hasPathSumDFS(root3, 1)
	common.PrintBool(r3, e3)

}
