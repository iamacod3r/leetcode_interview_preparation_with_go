package medium

import "interview_go/common"

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iii/description/
type LowestCommonAncestorOfaBinaryTreeIII struct {
}

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-tree-iii/solutions/950242/Multiple-solution-approaches-in-Java-(with-comments-and-explanation)/

// h1 is height of p
// h2 is height of q
// Time complexity is O(h1+h2)
// Space complexity is O(1)
func (l *LowestCommonAncestorOfaBinaryTreeIII) lowestCommonAncestor(p *Node, q *Node) *Node {
	a, b := p, q
	for a != b {
		a = a.Parent
		b = b.Parent
		if a == nil {
			a = q
		}

		if b == nil {
			b = p
		}
	}

	return b
}

func (l *LowestCommonAncestorOfaBinaryTreeIII) Test() {

	root := &Node{Val: 3}
	left := &Node{Parent: root, Val: 5}
	left.Left = &Node{Val: 6, Parent: left}
	left.Right = &Node{Parent: left, Val: 2}
	left.Right.Left = &Node{Parent: left.Right, Val: 7}
	left.Right.Right = &Node{Parent: left.Right, Val: 4}

	right := &Node{Parent: root, Val: 1}
	right.Left = &Node{Parent: right, Val: 0}
	right.Right = &Node{Parent: right, Val: 8}

	root.Left = left
	root.Right = right
	rs := l.lowestCommonAncestor(left.Right.Right, right.Left)
	e := 3
	if rs != nil {
		common.PrintInt(rs.Val, e)
	} else {
		println("result is nil")
	}
}
