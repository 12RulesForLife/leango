package main

import "fmt"

type Node struct {
	next *Node
	prev *Node
	val  int
}

type LinkedList struct {
	root *Node
	tail *Node
}

func (l *LinkedList) AddNode(val int) {
	if l.root == nil {
		l.root = &Node{val: val}
		l.tail = l.root
		return
	}

	l.tail.next = &Node{prev: l.tail, val: val}
	l.tail = l.tail.next
}

func (l *LinkedList) RemoveNode(node *Node) {
	if l.root == l.tail {
		l.root = nil
		l.tail = nil
		return
	}

	if node == l.root {
		l.root = l.root.next
		l.root.prev = nil
		return
	}

	if node == l.tail {
		l.tail = node.prev
		l.tail.next = nil
		return
	}

	node.next.prev = node.prev
	node.prev.next = node.next
}

func (l *LinkedList) PrintNodes() {
	if l.root == nil {
		fmt.Println()
		return
	}
	node := l.root
	for node != l.tail {
		fmt.Printf("%d - ", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)
}

func main() {
	var l LinkedList

	for i := 0; i < 10; i++ {
		l.AddNode(i)
		l.PrintNodes()
	}

	for i := 0; i < 10; i++ {
		l.RemoveNode(l.tail)
		l.PrintNodes()
	}

	for i := 0; i < 10; i++ {
		l.AddNode(i)
		l.PrintNodes()
	}

	for i := 0; i < 10; i++ {
		l.RemoveNode(l.root)
		l.PrintNodes()
	}

}
