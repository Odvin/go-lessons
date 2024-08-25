package main

import (
	"fmt"

	"github.com/Odvin/go-lessons/structures/list"
	"github.com/Odvin/go-lessons/structures/queue"
	"github.com/Odvin/go-lessons/structures/stack"
)

func main() {
	fmt.Println("--- Linkend List ---")

	linkedList := list.LinkedList{}

	nodeA := &list.Node{Data: 1}
	nodeB := &list.Node{Data: 2}
	nodeC := &list.Node{Data: 3}
	nodeD := &list.Node{Data: 4}
	nodeE := &list.Node{Data: 5}

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

	fmt.Println("--- Queue ---")

	queue := queue.Queue{}

	fmt.Println(queue)
	queue.Enqueue(100)
	fmt.Println(queue)
	queue.Enqueue(200)
	fmt.Println(queue)
	queue.Enqueue(300)
	fmt.Println(queue)

	fmt.Println(queue)
	qitem, err := queue.Dequeue()
	fmt.Println(queue, qitem, err)
	qitem, err = queue.Dequeue()
	fmt.Println(queue, qitem, err)
	qitem, err = queue.Dequeue()
	fmt.Println(queue, qitem, err)
	qitem, err = queue.Dequeue()
	fmt.Println(queue, qitem, err)
}