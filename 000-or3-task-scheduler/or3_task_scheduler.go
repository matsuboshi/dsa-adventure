package main 

import (
	"fmt"
)

func totalTimeToProcess(tm []int, tt []int, mm int) int {
	var totalTime int

	taskTypeMap := make(map[int][]int)

	for i := range tt {
		taskTypeMap[tt[i]] = append(taskTypeMap[tt[i]], tm[i]) 
	}

	for _, tmSlice := range taskTypeMap {
		if len(tmSlice) == 1 {
			totalTime++
			continue
		}

		if tmSlice[0] + tmSlice[1] > mm {
			totalTime += 2
		} else {
			totalTime++
		}
	}

	return totalTime
}

func main() {
	taskMemory := []int{1, 2, 3, 4, 2}
	taskType := []int{1, 2, 1, 2, 3}
	maxMemory := 4
	fmt.Println(totalTimeToProcess(taskMemory, taskType, maxMemory))

	taskMemory2 := []int{7, 2, 3, 9}
	taskType2 := []int{1, 2, 1, 3}
	maxMemory2 := 10
	fmt.Println(totalTimeToProcess(taskMemory2, taskType2, maxMemory2))
}
