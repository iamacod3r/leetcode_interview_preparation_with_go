package easy

import (
	"interview_go/common"
	"interview_go/datastructure"
	"math"
)

// https://leetcode.com/problems/maximum-depth-of-binary-tree/
type MaximumDepthOfBinaryTree struct {
}

func (m *MaximumDepthOfBinaryTree) maxDepth(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	}
	left := m.maxDepth(root.Left) + 1
	right := m.maxDepth(root.Right) + 1
	return int(math.Max(float64(left), float64(right)))
}

func (m *MaximumDepthOfBinaryTree) maxDepthWithIteration(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	}
	var stack []*datastructure.TreeNode
	var depths []float64
	stack = append(stack, root)
	depths = append(depths, 1)
	var depth, curr_depth float64 = 0, 0
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curr_depth = depths[len(depths)-1]
		depths = depths[:len(depths)-1]
		if node != nil {
			depth = math.Max(depth, curr_depth)
			if node.Left != nil {
				stack = append(stack, node.Left)
				depths = append(depths, curr_depth+1)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
				depths = append(depths, curr_depth+1)
			}
		}
	}
	return int(depth)
}

func (m *MaximumDepthOfBinaryTree) Test() {
	// [3,9,20,null,null,15,7]
	root := &datastructure.TreeNode{Val: 3}
	root.Left = &datastructure.TreeNode{Val: 9}
	rootRight := &datastructure.TreeNode{Val: 20}
	rootRight.Left = &datastructure.TreeNode{Val: 15}
	rootRight.Right = &datastructure.TreeNode{Val: 7}
	root.Right = rootRight
	e := 3
	r := m.maxDepthWithIteration(root)
	common.PrintInt(r, e)
}
