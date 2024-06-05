package main 

import "fmt"

func findMaxSubArray(nums []int, left int, right int) int {
	if left == right {
		return nums[left]
	} else {
		mid := (left+right)/2
		leftMax := findMaxSubArray(nums, left, mid)
		rightMax := findMaxSubArray(nums, mid+1, right)
		crossMax := maxCrossing(nums, left, right, mid)
		return max(leftMax, rightMax, crossMax)
	}
}

func maxCrossing(nums []int, left int, right int, mid int) int {
	leftSum := nums[mid]
	rightSum := nums[mid+1]

	sum :=0
	for i:=mid ; i >= left ; i-- {
		sum += nums[i]
		leftSum = max(leftSum, sum)
	}

	sum = 0
	for i:=mid+1 ; i <= right ; i++ {
		sum += nums[i]
		rightSum = max(rightSum, sum)
	}

	return leftSum + rightSum
}

func max(values ...int) int{
	maxVal := values[0]
	for _, v := range values {
		if maxVal<v {
			maxVal=v
		}
	}

	return maxVal
}

func main() {
	problem := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4} 
	answer := findMaxSubArray(problem,0,len(problem)-1)
	fmt.Println(answer)
}


