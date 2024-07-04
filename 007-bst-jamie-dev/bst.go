package main

import "fmt"

var count int

type Node struct {
	Key   int
	Left  *Node
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
	if n == nil {
		return false
	}

	if n.Key < k {
		return n.Right.Search(k)
	} else if n.Key > k {
		return n.Left.Search(k)
	}

	return true
}

func printNode(root *Node, k int) {
	if root != nil && count <= k {
		printNode(root.Right, k)
		count++
		if count == k {
			fmt.Println("GOTCHA:", root.Key)
		}
		printNode(root.Left, k)
	}
}

func main() {
	tree := Node{Key: 50}
	fmt.Println(tree)

	tree.Insert(30)
	tree.Insert(70)
	tree.Insert(20)
	tree.Insert(40)
	tree.Insert(60)
	tree.Insert(80)
	fmt.Println(tree)

	println(tree.Search(40))
	println(tree.Search(77))

	printNode(&tree, 3)
}
