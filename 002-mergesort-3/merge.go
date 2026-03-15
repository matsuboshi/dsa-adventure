// https://goplay.tools/snippet/ufex-vZIll7
package main

import (
	"fmt"
)

func mergeSort(array []int) []int {
	if len(array) < 2 {
		return array
	}

	mid := len(array) / 2
	left := mergeSort(array[:mid])
	right := mergeSort(array[mid:])

	return merge(left, right)
}

func merge(leftArr, rightArr []int) []int {
	cur, l, r := 0, 0, 0
	leftLen := len(leftArr)
	rightLen := len(rightArr)
	result := make([]int, leftLen+rightLen)

	for l < leftLen && r < rightLen {
		if leftArr[l] < rightArr[r] {
			result[cur] = leftArr[l]
			cur++
			l++
		} else {
			result[cur] = rightArr[r]
			cur++
			r++
		}
	}

	for l < leftLen {
		result[cur] = leftArr[l]
		cur++
		l++
	}

	for r < rightLen {
		result[cur] = rightArr[r]
		cur++
		r++
	}

	return result
}

func main() {
	testCases := [][]int{
		{1, 2, 7, 5},
		{10, 2, 3, 5},
		{10, 4, 2, 5, 1},
		{9, 6, 5, 6, 7, 1, 4, 5, 3, 3, 8, 9, 2, 6, 3, 6, 1},
	}

	for _, tc := range testCases {
		fmt.Println(mergeSort(tc))
	}

}
