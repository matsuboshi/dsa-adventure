package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	Data T
	height int
	leftChild *Node[T]
	rightChild *Node[T]
}

func (n *Node[T]) balanceFactor() int {
	if n==nil {
		return 0
	}

	leftHeight := 0
	if n.leftChild != nil {
		leftHeight = n.leftChild.height
	}

	rightHeight := 0
	if n.rightChild != nil {
		rightHeight = n.rightChild.height
	}

	return leftHeight - rightHeight
}

func height[T constraints.Ordered](node *Node[T]) int {
	height := 0
	if node != nil {
		height = node.height
	}

	return height
}

func rotateLeft[T constraints.Ordered](node *Node[T]) *Node[T] {
	y := node.rightChild
	tmp := y.leftChild 

	y.leftChild = node
	node.rightChild = tmp

	node.height = max(
		height(node.leftChild),
		height(node.rightChild),
	) + 1

	y.height = max(
		height(y.leftChild),
		height(y.rightChild),
	) + 1

	return y
}

func rotateRight[T constraints.Ordered](node *Node[T]) *Node[T] {
	x := node.leftChild
	tmp := x.rightChild 

	x.rightChild = node
	node.leftChild = tmp

	node.height = max(
		height(node.leftChild),
		height(node.rightChild),
	) + 1

	x.height = max(
		height(x.leftChild),
		height(x.rightChild),
	) + 1

	return x
}

func insert[T constraints.Ordered](node *Node[T], data T) *Node[T] {
	if node == nil {
		node = &Node[T]{Data: data}
		return node
	}

	if data < node.Data {
		node.leftChild = insert(node.leftChild, data)
	} else if data > node.Data {
		node.rightChild = insert(node.rightChild, data)
	} else {
		return node
	}

	node.height = 1 + max(
		height(node.leftChild),
		height(node.rightChild),
	)

	bf := node.balanceFactor()

	if bf > 1 && data < node.leftChild.Data {
		return rotateRight(node)
	}

	if bf < -1 && data > node.rightChild.Data {
		return rotateLeft(node)
	}

	if bf > 1 && data > node.leftChild.Data {
		node.leftChild = rotateLeft(node.leftChild)
		return rotateRight(node)
	}

	if bf < -1 && data < node.rightChild.Data {
		node.rightChild = rotateRight(node.rightChild)
		return rotateLeft(node)
	}

	return node
}

func inOrderTraversalData[T constraints.Ordered](node *Node[T], vals *[]T) {
	if node != nil {
		inOrderTraversalData(node.leftChild, vals)
		*vals = append(*vals, node.Data)
		inOrderTraversalData(node.rightChild, vals)
	}
}

func inOrderTraversalHeight[T constraints.Ordered](node *Node[T], vals *[]int) {
	if node != nil {
		inOrderTraversalHeight(node.leftChild, vals)
		*vals = append(*vals, node.height)
		inOrderTraversalHeight(node.rightChild, vals)
	}
}

func printAVLTree[T constraints.Ordered](root *Node[T]) {
	dataVals := []T{}
	inOrderTraversalData(root, &dataVals)

	heightVals := []int{}
	inOrderTraversalHeight(root, &heightVals)

	fmt.Println(dataVals)
	fmt.Println(heightVals)
	println()

	cols := len(heightVals)
	rows := 0
	for _, v := range heightVals {
		curr := v+1
		if rows < curr { rows = curr }
	}

	
	matrix := make([][]string, rows)
	for i := range matrix {
		matrix[i] = make([]string, cols)
	}

	for i := range matrix {
		for j := range matrix[i] {
			matrix[i][j] = fmt.Sprintf("%2s", "")
		}
	}

	for i, v := range heightVals {
		matrix[v][i] = fmt.Sprintf("%2v", dataVals[i])
	}

	for i := range matrix {
		fmt.Println(matrix[rows-1-i])
	}

}

func main() {
	root := &Node[int]{Data: 6}
	fmt.Println(root.Data)
	root = insert(root, 27)
	root = insert(root, 38)
	root = insert(root, 29)
	root = insert(root, 15)
	root = insert(root, 34)
	root = insert(root, 22)
	root = insert(root, 12)
	root = insert(root, 36)
	root = insert(root, 30)
	
	printAVLTree(root)
}

