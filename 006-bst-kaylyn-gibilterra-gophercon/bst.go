package main

import "fmt"

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

func (n *Node) Search(key int) bool {
	if n==nil { return false }

	if n.Key < key {
		return n.Right.Search(key)
	} else if n.Key > key {
		return n.Left.Search(key)  
	}

	return true
}

func (n *Node) Insert(key int) {
	if n.Key < key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	} else if n.Key > key {
		 if n.Left == nil {
			n.Left = &Node{Key: key}
		 } else {
			n.Left.Insert(key)
		 }
	}
}

func (n *Node) Delete(key int) *Node {
	if n.Key < key {
		n.Right = n.Right.Delete(key)
	} else if n.Key > key {
		n.Left = n.Left.Delete(key)
	} else {
		if n.Left == nil {
			return n.Right
		} else if n.Right == nil {
			return n.Left
		}

		min := n.Right.Min()
		n.Key = min 
		n.Right = n.Right.Delete(min)
	}
	return n
}

func (n *Node) Min() int {
	if n.Left == nil {
		return n.Key
	}
	return n.Left.Min()
}

// func (n *Node) Max() int {
// 	if n.Right == nil {
// 		return n.Key
// 	}
// 	return n.Right.Max()
// }

func main() {
	tree := &Node{6, nil, nil}
	fmt.Println(tree.Key)
	fmt.Println("ROOT", tree.Search(6))

	fmt.Println(tree.Search(7))
	fmt.Println(tree.Search(4))
	fmt.Println(tree.Search(2))
	fmt.Println(tree.Search(9))

	tree.Insert(7)
	tree.Insert(4)
	tree.Insert(2)
	tree.Insert(9)

	println()
	fmt.Println(tree.Search(7))
	fmt.Println(tree.Search(4))
	fmt.Println(tree.Search(2))
	fmt.Println(tree.Search(9))
	
	println()
	tree.Delete(7)
	fmt.Println(tree.Search(7))
	fmt.Println(tree.Search(4))
	fmt.Println(tree.Search(2))
	fmt.Println(tree.Search(9))
	
	println()
	tree.Delete(4)
	fmt.Println(tree.Search(7))
	fmt.Println(tree.Search(4))
	fmt.Println(tree.Search(2))
	fmt.Println(tree.Search(9))

	println()
	tree.Delete(2)
	fmt.Println(tree.Search(7))
	fmt.Println(tree.Search(4))
	fmt.Println(tree.Search(2))
	fmt.Println(tree.Search(9))

	println()
	tree.Delete(9)
	fmt.Println(tree.Search(7))
	fmt.Println(tree.Search(4))
	fmt.Println(tree.Search(2))
	fmt.Println(tree.Search(9))
}

