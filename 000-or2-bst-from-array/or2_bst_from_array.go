package main 

import (
	"fmt"
)

func left(i int) int { return (2*i)+1 }

func right(i int) int { return (2*i)+2 }

func findNumber (tree []int, index int, target int) int {
	if index >= len(tree) {
		return 0
	}

	if target < tree[index] {
		return findNumber(tree, left(index), target)
	}

	if target > tree[index] {
		return findNumber(tree, right(index), target)
	}

	return 1
}

func main() {
	list := []int{ 30, 10, 12, 15 }

	tree := []int{ 20, 10, 30, 8, 12, 25, 40, 6, 11, 13, 23 }

	for _, v := range list {
		fmt.Println(v, findNumber(tree, 0, v))
	}
}

