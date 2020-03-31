package main

import "fmt"

type Node struct {
	next *Node
	val  int
}

func main() {
	var root *Node
	root = &Node{val: 0}
	tail := root

	PrintNode(root)

	root, tail = RemoveNode(tail, root, tail)
	PrintNode(root)
	fmt.Print(root, tail)
}

func AddNode(tail *Node, val int) *Node {
	node := &Node{val: val}
	tail.next = node
	return node
}

func RemoveNode(node *Node, root *Node, tail *Node) (*Node, *Node) {
	if node == root {
		root = root.next

		if root == nil {
			tail = nil
		}
		return root, tail

	}

	prev := root
	for prev.next != node {
		prev = prev.next
	}
	if node == tail {
		tail = prev
		prev.next = nil
		return root, tail
	}

	prev.next = node.next
	return root, tail
}

func PrintNode(node *Node) {
	if node == nil {
		return
	}
	for node.next != nil {
		fmt.Printf("%d ->", node.val)
		node = node.next
	}
	fmt.Printf("%d\n", node.val)

}
