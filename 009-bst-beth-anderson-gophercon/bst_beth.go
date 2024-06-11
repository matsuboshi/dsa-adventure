package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	Value T
	Left *Node[T]
	Right *Node[T]
}

type BinaryTree[T constraints.Ordered] struct {
	Root *Node[T]
	Counter int
}

func (bt *BinaryTree[T]) Add(o T) {
	if bt.Root == nil {
		bt.Root = &Node[T]{
			Value: o,
		}
		bt.Counter ++
	} else {
		bt.addTo(bt.Root, o)
	}
} 

func (bt *BinaryTree[T]) addTo(node *Node[T], val T) {
	nd := Node[T]{
		Value: val,
	}

	if nd.Value < node.Value {
		if node.Left == nil {
			node.Left = &nd
			bt.Counter ++
		} else {
			bt.addTo(node.Left, val)
		}
	} else if nd.Value > node.Value {
		if node.Right == nil {
			node.Right = &nd
			bt.Counter ++
		} else {
			bt.addTo(node.Right, val)
		}
	}
}

func (bt *BinaryTree[T]) inOrderTraversal(node *Node[T], vals *[]T) {
	if node != nil {
		bt.inOrderTraversal(node.Left, vals)
		*vals = append(*vals, node.Value)
		bt.inOrderTraversal(node.Right, vals)
	}
}

func main() {
	bt := &BinaryTree[int]{}
	bt.Add(7)
	bt.Add(6)
	bt.Add(9)
	bt.Add(8)
	bt.Add(2)
	bt.Add(3)
	bt.Add(4)
	bt.Add(5)
	println(bt.Counter)

	vals := []int{}
	bt.inOrderTraversal(bt.Root, &vals)
	fmt.Println(vals)
}
