package tree

type Node struct {
	Key   int
	left  *Node
	right *Node
}

func (n *Node) Insert(key int) {
	if n.Key < key {
		// to right
		if n.right == nil {
			n.right = &Node{Key: key}
		} else {
			n.right.Insert(key)
		}
	} else {
		// to left
		if n.left == nil {
			n.left = &Node{Key: key}
		} else {
			n.left.Insert(key)
		}
	}
}

func (n *Node) Search(key int) bool {
	if n == nil {
		return false
	}

	if n.Key == key {
		return true
	}

	if n.Key < key {
		// to right
		return n.right.Search(key)
	} else {
		// to left
		return n.left.Search(key)
	}
}
