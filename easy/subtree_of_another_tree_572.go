package easy

import (
	"interview_go/common"
	"interview_go/datastructure"
	"strings"
)

// https://leetcode.com/problems/subtree-of-another-tree/description/
type SubtreeOfAnotherTree struct{}

// M is length of root
// N is length of subRoot
// Time complexity: O(MN)
// Space complexity: O(M+N)
func (s *SubtreeOfAnotherTree) isSubtree(root *datastructure.TreeNode, subRoot *datastructure.TreeNode) bool {

	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	}

	if s.isIdentical(root, subRoot) {
		return true
	}

	return s.isSubtree(root.Left, subRoot) || s.isSubtree(root.Right, subRoot)
}

func (s *SubtreeOfAnotherTree) isIdentical(root, subRoot *datastructure.TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	}

	return root.Val == subRoot.Val && s.isIdentical(root.Left, subRoot.Left) && s.isIdentical(root.Right, subRoot.Right)
}

// String Matching Algorithms with Serialize

func (s *SubtreeOfAnotherTree) serialize(node *datastructure.TreeNode, sb *strings.Builder) {
	if node == nil {
		sb.WriteString("#")
		return
	}
	sb.WriteString("^")
	sb.WriteRune(rune(node.Val))
	s.serialize(node.Left, sb)
	s.serialize(node.Right, sb)
}

func (s *SubtreeOfAnotherTree) kmp(haystack, needle string) bool {
	m := len(needle)
	n := len(haystack)

	if m > n {
		return false
	}

	lps := make([]int, m)
	prev, i := 0, 1

	for i < m {
		if needle[i] == needle[prev] {
			prev++
			lps[i] = prev
			i++
		} else {
			if prev == 0 {
				lps[i] = 0
				i++
			} else {
				prev = lps[prev-1]
			}
		}
	}

	hp, np := 0, 0

	for hp < n {
		if haystack[hp] == needle[np] {
			np++
			hp++
			if np == m {
				return true
			}
		} else {
			if np == 0 {
				// zero matched
				hp++
			} else {
				np = lps[np-1]
			}
		}
	}

	return false
}

func (s *SubtreeOfAnotherTree) isSubtreeSerialize(root, subRoot *datastructure.TreeNode) bool {

	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	}

	var sbr strings.Builder
	s.serialize(root, &sbr)
	rtr := sbr.String()

	var sbs strings.Builder
	s.serialize(subRoot, &sbs)
	str := sbs.String()

	return s.kmp(rtr, str)
}

const (
	mod_1 = 1000000007
	mod_2 = 2147483647
)

var memo [][]int64

func (s *SubtreeOfAnotherTree) hahsSubtreeAtNode(node *datastructure.TreeNode, needToAdd bool) []int64 {
	if node == nil {
		return []int64{3, 7}
	}

	leftNode := s.hahsSubtreeAtNode(node.Left, needToAdd)
	rightNode := s.hahsSubtreeAtNode(node.Right, needToAdd)

	left1 := (leftNode[0] << 5) % mod_1
	right1 := (rightNode[0] << 1) % mod_1
	left2 := (leftNode[1] << 7) % mod_2
	right2 := (rightNode[1] << 1) % mod_2

	hashPair := []int64{
		(left1 + right1 + int64(node.Val)) % mod_1,
		(left2 + right2 + int64(node.Val)) % mod_2,
	}

	if needToAdd {
		memo = append(memo, hashPair)
	}
	return hashPair
}

func (s *SubtreeOfAnotherTree) isSubtreeWithHashing(root, subRoot *datastructure.TreeNode) bool {
	if root == nil || subRoot == nil {
		return root == nil && subRoot == nil
	}

	s.hahsSubtreeAtNode(root, true)

	sub := s.hahsSubtreeAtNode(subRoot, false)

	common.Print2Dint64(memo, [][]int64{sub})

	for _, m := range memo {

		if m[0] == sub[0] && m[1] == sub[1] {
			return true
		}
	}
	return false
}

func (s *SubtreeOfAnotherTree) Test() {
	subRoot := &datastructure.TreeNode{Val: 4, Left: &datastructure.TreeNode{Val: 1}, Right: &datastructure.TreeNode{Val: 2}}
	root := &datastructure.TreeNode{Val: 3, Left: subRoot, Right: &datastructure.TreeNode{Val: 5}}

	e := true
	r := s.isSubtreeWithHashing(root, subRoot)
	common.PrintBool(r, e)
}
