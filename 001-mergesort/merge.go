package main

import (
	"fmt"
)

func mergeSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	halfSize := len(numbers) / 2

	left := mergeSort(numbers[:halfSize])
	right := mergeSort(numbers[halfSize:])

	return merge(left, right)
}



func merge(left []int, right []int) []int {
	li := 0
	ri := 0
	
	
	

	fullSize := len(left) + len(right)
	result := make([]int, fullSize)

	for i := 0; i < fullSize; i++ {
		if li < len(left) && ri < len(right) {
			if left[li] < right[ri] {
				result[i] = left[li]
				li++
			} else {
				result[i] = right[ri]
				ri++
			}
			continue
		}
		if li < len(left) {
			result[i] = left[li]
			li++
			continue
		}
		if ri < len(right) {
			result[i] = right[ri]
			ri++
			continue
		}
	}

	return result
}

func main() {
	sorted := mergeSort([]int{1, 2, 3, 7, 4, 3, 7, 3, 2, 9, 3, 6, 7})
	fmt.Println(sorted)
}
