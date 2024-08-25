package main

import (
	"fmt"

	linkedlisl "github.com/Odvin/go-lessons/structures/linked_list"
	"github.com/Odvin/go-lessons/structures/stack"
)

func main() {
	fmt.Println("--- Linkend List ---")

	linkedList := linkedlisl.LinkedList{}

	nodeA := &linkedlisl.Node{Data: 1}
	nodeB := &linkedlisl.Node{Data: 2}
	nodeC := &linkedlisl.Node{Data: 3}
	nodeD := &linkedlisl.Node{Data: 4}
	nodeE := &linkedlisl.Node{Data: 5}

	linkedList.Prepend(nodeA)
	linkedList.Prepend(nodeB)
	linkedList.Prepend(nodeC)
	linkedList.Prepend(nodeD)
	linkedList.Prepend(nodeE)

	linkedList.Print()
	fmt.Println(linkedList.HasData(3))
	linkedList.DelValue(3)
	fmt.Println(linkedList.HasData(3))
	linkedList.Print()

	fmt.Println("--- Stack ---")

	stack := stack.Stack{}

	stack.Push(100)
	stack.Push(200)
	stack.Push(300)

	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
	stack.Pop()
	fmt.Println(stack)
}
