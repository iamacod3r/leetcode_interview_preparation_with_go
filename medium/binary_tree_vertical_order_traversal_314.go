package medium

import (
	"interview_go/common"
	"interview_go/datastructure"
	"sort"
)

// https://leetcode.com/problems/binary-tree-vertical-order-traversal/
type BinaryTreeVerticalOrderTraversal struct {
}

type PairStruct struct {
	TreeNode *datastructure.TreeNode
	Column   int
}

// ****** FASTER ******
// Time Complexity : O(N) - no sorting
// Space Complexity : O(N)
func (t *BinaryTreeVerticalOrderTraversal) verticalOrder_with_Modified_BFS(root *datastructure.TreeNode) [][]int {
	result := [][]int{}
	if root == nil {
		return result
	}

	colTable := map[int][]int{}
	queue := []PairStruct{{root, 0}}
	minColumn, maxColumn := 0, 0

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		node := p.TreeNode
		if node != nil {
			if _, ok := colTable[p.Column]; !ok {
				colTable[p.Column] = []int{}
			}
			colTable[p.Column] = append(colTable[p.Column], node.Val)
			minColumn = common.Min(minColumn, p.Column)
			maxColumn = common.Max(maxColumn, p.Column)
			queue = append(queue, PairStruct{node.Left, p.Column - 1})
			queue = append(queue, PairStruct{node.Right, p.Column + 1})
		}
	}

	for i := minColumn; i <= maxColumn; i++ {
		result = append(result, colTable[i])
	}

	return result
}

func (t *BinaryTreeVerticalOrderTraversal) verticalOrder_BFS2(root *datastructure.TreeNode) [][]int {
	println("verticalOrder_BFS2")

	result := [][]int{}
	if root == nil {
		return result
	}

	colTable := map[int][]int{}
	queue := []PairStruct{{root, 0}}
	minCol, maxCol := 0, 0

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		node := p.TreeNode
		if node != nil {

			if _, ok := colTable[p.Column]; !ok {
				colTable[p.Column] = []int{node.Val}
			} else {
				colTable[p.Column] = append(colTable[p.Column], node.Val)
			}

			minCol = common.Min(minCol, p.Column)
			maxCol = common.Max(maxCol, p.Column)
			if node.Left != nil {
				queue = append(queue, PairStruct{node.Left, p.Column - 1})
			}
			if node.Right != nil {
				queue = append(queue, PairStruct{node.Right, p.Column + 1})
			}
		}
	}

	for i := minCol; i <= maxCol; i++ {
		result = append(result, colTable[i])
	}

	return result
}

// Time Complexity : O(nlogN) - because of Sorting
// Space Complexity : O(N)
func (t *BinaryTreeVerticalOrderTraversal) verticalOrderBFS(root *datastructure.TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	columnTable := make(map[int][]int)
	queue := []PairStruct{{root, 0}}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		if p.TreeNode != nil {
			if _, ok := columnTable[p.Column]; !ok {
				columnTable[p.Column] = []int{}
			}

			columnTable[p.Column] = append(columnTable[p.Column], p.TreeNode.Val)

			if p.TreeNode.Left != nil {
				queue = append(queue, PairStruct{p.TreeNode.Left, p.Column - 1})
			}

			if p.TreeNode.Right != nil {
				queue = append(queue, PairStruct{p.TreeNode.Right, p.Column + 1})
			}
		}
	}

	keys := []int{}
	for k, _ := range columnTable {
		keys = append(keys, k)
	}

	sort.Ints(keys) // O(nlogN)

	for _, k := range keys {
		result = append(result, columnTable[k])
	}

	return result
}

var columnsTable map[int][][]int
var minCol, maxCol int = 0, 0

// N is the number of nodes in the tree
// W is the width of the binary tree
// H is the height of the binary tree
// Time Complexity : O(W * H log H)
// Space Complexity : O(n)
func (t *BinaryTreeVerticalOrderTraversal) verticalOrderDFS(root *datastructure.TreeNode) [][]int {
	result := [][]int{}

	if root == nil {
		return result
	}

	columnsTable = map[int][][]int{}
	minCol = 0
	maxCol = 0

	t.DFS(root, 0, 0)

	for i := minCol; i < maxCol+1; i++ {
		// sort

		sort.SliceStable(columnsTable[i], func(l, b int) bool {
			return columnsTable[i][l][0] < columnsTable[i][b][0]
		})

		sortedColumn := []int{}

		for _, v := range columnsTable[i] {
			sortedColumn = append(sortedColumn, v[1])
		}

		result = append(result, sortedColumn)
	}

	return result
}

func (t *BinaryTreeVerticalOrderTraversal) DFS(node *datastructure.TreeNode, row, col int) {
	if node == nil {
		return
	}

	if _, ok := columnsTable[col]; !ok {
		columnsTable[col] = [][]int{{row, node.Val}}
	} else {
		columnsTable[col] = append(columnsTable[col], []int{row, node.Val})
	}

	minCol = common.Min(minCol, col)
	maxCol = common.Max(maxCol, col)

	t.DFS(node.Left, row+1, col-1)
	t.DFS(node.Right, row+1, col+1)
}

func (t *BinaryTreeVerticalOrderTraversal) Test() {
	right := &datastructure.TreeNode{Val: 20, Left: &datastructure.TreeNode{Val: 15}, Right: &datastructure.TreeNode{Val: 7}}
	root := &datastructure.TreeNode{Val: 3, Left: &datastructure.TreeNode{Val: 9}, Right: right}
	e := [][]int{
		{9},
		{3, 15},
		{20},
		{7},
	}
	r := t.verticalOrder_BFS2(root)
	common.Print2DSlice(r, e)

	l2 := &datastructure.TreeNode{Val: 9, Left: &datastructure.TreeNode{Val: 4}, Right: &datastructure.TreeNode{Val: 0}}
	r2 := &datastructure.TreeNode{Val: 8, Left: &datastructure.TreeNode{Val: 1}, Right: &datastructure.TreeNode{Val: 7}}
	root2 := &datastructure.TreeNode{Val: 3, Left: l2, Right: r2}
	e2 := [][]int{
		{4},
		{9},
		{3, 0, 1},
		{8},
		{7},
	}
	res2 := t.verticalOrder_BFS2(root2)
	common.Print2DSlice(res2, e2)

	l3 := &datastructure.TreeNode{Val: 9, Left: &datastructure.TreeNode{Val: 4}, Right: &datastructure.TreeNode{Val: 0, Left: &datastructure.TreeNode{Val: 5}}}
	r3 := &datastructure.TreeNode{Val: 8, Left: &datastructure.TreeNode{Val: 1, Right: &datastructure.TreeNode{Val: 2}}, Right: &datastructure.TreeNode{Val: 7}}
	root3 := &datastructure.TreeNode{Val: 3, Left: l3, Right: r3}
	e3 := [][]int{
		{4},
		{9, 5},
		{3, 0, 1},
		{8, 2},
		{7},
	}
	res3 := t.verticalOrder_BFS2(root3)
	common.Print2DSlice(res3, e3)

}
