package list

import "fmt"

type Node struct {
	Data int
	next *Node
}

type LinkedList struct {
	head   *Node
	length int
}

func (l *LinkedList) Prepend(n *Node) {
	node := l.head
	l.head = n
	l.head.next = node
	l.length++
}

func (l *LinkedList) Print() {
	node := l.head
	for i := 0; i < l.length; i++ {
		fmt.Printf("%d ", node.Data)
		node = node.next
	}
	fmt.Printf("\n")
}

func (l *LinkedList) DelValue(data int) {

	for l.head.Data == data {
		l.head = l.head.next
		l.length--
	}

	node := l.head
	nextNode := l.head.next
	nodes := l.length

	for i := 1; i < nodes; i++ {
		if nextNode.Data == data {
			node.next = nextNode.next
			nextNode = nextNode.next
			l.length--
		} else {
			node = nextNode
			nextNode = nextNode.next
		}
	}
}

func (l *LinkedList) HasData(data int) bool {
	node := l.head

	for i := 0; i < l.length; i++ {
		if node.Data == data {
			return true
		} else {
			node = node.next
		}
	}

	return false
}
