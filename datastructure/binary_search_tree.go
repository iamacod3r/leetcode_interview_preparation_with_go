package datastructure

import "fmt"

type BSTNode struct {
	value int
	left  *BSTNode
	right *BSTNode
}

func (n *BSTNode) Insert(val int) {
	if n.value < val {
		// move right
		if n.right == nil {
			n.right = &BSTNode{value: val}
		} else {
			n.right.Insert(val)
		}
	} else if n.value > val {
		// move left
		if n.left == nil {
			n.left = &BSTNode{value: val}
		} else {
			n.left.Insert(val)
		}
	}
}

func (n *BSTNode) Search(val int) (*BSTNode, error) {
	if n == nil {
		return nil, fmt.Errorf("value can not find")
	}
	if n.value < val {
		return n.right.Search(val)
	}
	if n.value > val {
		return n.left.Search(val)
	}

	return n, nil
}

type BinarySearchTree struct {
	root *BSTNode
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}

func (b *BinarySearchTree) Reset() {
	b.root = nil
}

func (b *BinarySearchTree) InsertAlternative(value int) {

}

func (b *BinarySearchTree) Insert(value int) {
	b.InsertValue(b.root, value)
}

func (b *BinarySearchTree) InsertValue(node *BSTNode, value int) {

	if b.root == nil {
		b.root = &BSTNode{
			value: value,
		}
		return
	}

	if node == nil {
		return
	}

	var tempNode *BSTNode
	for node != nil {
		tempNode = node
		if value < node.value {
			node = node.left
		} else {
			node = node.right
		}
	}
	if value < tempNode.value {
		tempNode.left = &BSTNode{value: value}
	} else {
		tempNode.right = &BSTNode{value: value}
	}
}

func (b *BinarySearchTree) Find(value int) error {
	node := b.FindValue(b.root, value)
	if node == nil {
		return fmt.Errorf("find error : %d not found in tree", value)
	}
	return nil
}

func (b *BinarySearchTree) FindValue(node *BSTNode, value int) *BSTNode {
	if node == nil {
		return nil
	}
	if node.value == value {
		return node
	}
	if value < node.value {
		return b.FindValue(node.left, value)
	}
	return b.FindValue(node.right, value)
}

func (b *BinarySearchTree) InOrder() {
	b.InOrderWithNode(b.root)
}

func (b *BinarySearchTree) InOrderWithNode(node *BSTNode) {
	if node != nil {
		b.InOrderWithNode(node.left)
		fmt.Println(node.value)
		b.InOrderWithNode(node.right)
	}
}
