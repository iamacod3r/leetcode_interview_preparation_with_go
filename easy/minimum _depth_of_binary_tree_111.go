package easy

import (
	"interview_go/common"
	"interview_go/datastructure"
	"math"
)

// https://leetcode.com/problems/minimum-depth-of-binary-tree/
type MinimumDepthOfBinaryTree struct {
}

/*
 - BFS will stop the moment we reach the level of min-depth. DFS will continue searching all the way down.On average, for randomly generated trees, the min-depth will be much smaller than the max-depth, so BFS will access fewer nodes.
 - BFS's queue takes less memory than a call stack which has to remember the functions' params local variables, return address etc. This can lead to a stack overflow error as memory consumption increases beyond its limits.
 - DFS is calling the function recursively, and popular languages like Python, JavaScript have a max recursion depth of 1000. Hence, it cannot solve a tree with 1001 right nodes.
*/

func (m *MinimumDepthOfBinaryTree) minDepthDFS(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	}

	if root.Left == nil {
		return m.minDepthDFS(root.Right) + 1
	} else if root.Right == nil {
		return m.minDepthDFS(root.Left) + 1
	}

	left := m.minDepthDFS(root.Left) + 1
	right := m.minDepthDFS(root.Right) + 1
	return int(math.Min(float64(left), float64(right)))
}

func (m *MinimumDepthOfBinaryTree) minDepthBFS(root *datastructure.TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	q := []*datastructure.TreeNode{root}

	for len(q) > 0 {
		curr := q[0]
		lenQ := len(q)
		q = q[:lenQ-1]
		println(len(q))
		if curr == nil {
			continue
		}
		if curr.Left == nil && curr.Right == nil {
			return depth
		}
		if curr.Left != nil {
			q = append(q, curr.Left)
		}
		if curr.Right != nil {
			println("curr.Right", len(q))
			q = append(q, curr.Right)
		}
		depth++
	}

	return -1
}

func (m *MinimumDepthOfBinaryTree) Test() {

	root := &datastructure.TreeNode{Val: 3}
	root.Left = &datastructure.TreeNode{Val: 9}
	rootRight := &datastructure.TreeNode{Val: 20}
	rootRight.Left = &datastructure.TreeNode{Val: 15}
	rootRight.Right = &datastructure.TreeNode{Val: 7}
	root.Right = rootRight
	e := 2
	r := m.minDepthBFS(root)
	common.PrintInt(r, e)

	right5 := &datastructure.TreeNode{Val: 6}
	right4 := &datastructure.TreeNode{Val: 5, Right: right5}
	right3 := &datastructure.TreeNode{Val: 4, Right: right4}
	right2 := &datastructure.TreeNode{Val: 3, Right: right3}
	root2 := &datastructure.TreeNode{Val: 2, Right: right2}
	e2 := 5
	r2 := m.minDepthBFS(root2)
	common.PrintInt(r2, e2)

}
