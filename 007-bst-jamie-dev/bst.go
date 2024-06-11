package main

import "fmt"

type Node struct {
	Key int
	Left *Node
	Right *Node
}

func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	if n == nil { return false }

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true
} 



func main() {
	tree := Node{Key: 100}
	fmt.Println(tree)

	tree.Insert(52)
	tree.Insert(203)
	tree.Insert(19)
	tree.Insert(76)
	tree.Insert(310)
	tree.Insert(88)
	tree.Insert(276)
	fmt.Println(tree)

	println(tree.Search(34))
	println(tree.Search(77))
	println(tree.Search(99))
	println(tree.Search(88))
	println(tree.Search(76))
	println(tree.Search(19))
}

