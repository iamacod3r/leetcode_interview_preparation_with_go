package datastructure

const ALPHABET_SIZE = 26

type TrieNode1 struct {
	children [ALPHABET_SIZE]*TrieNode1
	IsWord   bool
}

type Trie1 struct {
	root *TrieNode1
}

func NewTrie1() *Trie1 {
	return &Trie1{
		root: &TrieNode1{},
	}
}

func (t *Trie1) Insert(word string) {
	wordLen := len(word)
	current := t.root
	for i := 0; i < wordLen; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			current.children[index] = &TrieNode1{}
		}
		current = current.children[index]
	}
	current.IsWord = true
}

func (t *Trie1) Find(word string) bool {
	wordLen := len(word)
	current := t.root
	for i := 0; i < wordLen; i++ {
		index := word[i] - 'a'
		if current.children[index] == nil {
			return false
		}
		current = current.children[index]
	}

	return current.IsWord
}
