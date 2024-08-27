package trie

// Number of the possible characters in the trie
const AlphabetSize = 26

// Node in the trie
type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

// Trie
type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	return &Trie{
		root: &Node{},
	}
}

// Insert
func (t *Trie) Insert(word string) {
	wordLength := len(word)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			// create node
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Search
func (t *Trie) Search(word string) bool {
	wordLength := len(word)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := word[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		} else {
			currentNode = currentNode.children[charIndex]
		}
	}

	return currentNode.isEnd
}
