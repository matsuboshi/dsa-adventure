package main

import "fmt"

func MergeSort(nums []int) []int {
	// base case
	if len(nums) < 2 {
		return nums
	}

	halfSize := len(nums)/2
	left := MergeSort(nums[:halfSize])
	right := MergeSort(nums[halfSize:])

	return merge(left, right)
}

func merge (left, right []int) []int {
	li := 0
	ri := 0

	leftSize := len(left)
	rightSize := len(right)
	fullSize := leftSize + rightSize
	result := make([]int, fullSize)

	for i:=0 ; i<fullSize ; i++ { 
		if li<leftSize && ri<rightSize {
			if left[li]<right[ri] {
				result[i] = left[li]
				li++
			} else {
				result[i] = right[ri]
				ri++
			}
			continue
		}

		if li<leftSize {
			result[i] = left[li]
			li++
			continue
		}

		if ri<rightSize {
			result[i] = right[ri]
			ri++
			continue
		}
	}

	return result
}


func main() {
	sortedSlice := MergeSort([]int{9,6,5,6,7,1,4,5,3,3,8,9,2,6,3,6,1})

	fmt.Println(sortedSlice)
}