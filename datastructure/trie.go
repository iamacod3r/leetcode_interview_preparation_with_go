package datastructure

type TrieNode struct {
	children map[rune]*TrieNode
	isWord   bool
}

type Trie struct {
	root *TrieNode
}

func (t *Trie) Insert(w string) {

	curr := t.root
	for _, letter := range w {
		child, ok := curr.children[letter]
		if !ok {
			child = &TrieNode{
				children: make(map[rune]*TrieNode),
				isWord:   false,
			}
			curr.children[letter] = child
		}
		curr = child
	}
	curr.isWord = true
}

func (t *Trie) Search(w string) bool {

	curr := t.root
	for _, letter := range w {
		child, ok := curr.children[letter]
		if !ok {
			return false
		}
		curr = child
	}
	return curr.isWord
}

func InitTrie() *Trie {
	return &Trie{root: &TrieNode{
		children: make(map[rune]*TrieNode),
		isWord:   false,
	}}
}
