package medium

type LowestCommonAncestorOfaBinaryTreeIII struct {
}

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func (l *LowestCommonAncestorOfaBinaryTreeIII) lowestCommonAncestor(p *Node, q *Node) *Node {
	return &Node{}
}

func (l *LowestCommonAncestorOfaBinaryTreeIII) Test() {

	lr := &Node{Val: 2, Left: &Node{Val: 7}, Right: &Node{Val: 4}}
	left := &Node{Val: 5, Left: &Node{Val: 6}, Right: lr}
	right := &Node{Val: 1, Left: &Node{Val: 0}, Right: &Node{Val: 8}}
	root := &Node{Val: 3, Left: left, Right: right}
	l.lowestCommonAncestor(root, left)

}
